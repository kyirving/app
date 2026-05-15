package main

import (
	"app/config"
	"app/internal/app"
	"app/internal/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	args := &config.CommondArgs{}
	pflag.StringVar(&args.ConfigDir, "configDir", "./config", "配置文件目录")
	pflag.StringVar(&args.ConfigFile, "configFile", "config.yaml", "配置文件")
	pflag.Parse()
	// 解析命令行参数
	cfg, err := config.LoadConfig(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 连接数据库
	db, err := database.ConnectMySQL(&cfg.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	server := app.NewServer(cfg, db)

	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-quit
	fmt.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}

	// 关闭数据库连接
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
	}

	fmt.Println("server stopped")
}
