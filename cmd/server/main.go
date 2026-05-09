package main

import (
	"app/config"
	"app/internal/app"
	"app/internal/database"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/pflag"
)

func main() {
	args := &config.CommondArgs{}
	pflag.StringVar(&args.ConfigDir, "config-dir", "/Users/wuh/data/goland/app/config", "配置文件目录")
	pflag.StringVar(&args.ConfigFile, "config-file", "config.yaml", "配置文件")
	pflag.Parse()
	// 解析命令行参数
	config, err := config.LoadConfig(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 连接数据库
	db, err := database.ConnectMySQL(&config.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 优雅退出
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		<-signals
		// 关闭数据库连接
		if DB, err := db.DB(); err == nil {
			DB.Close()
		}

		// 关闭Web服务
		fmt.Println("web服务关闭")
		os.Exit(0)
	}()

	server := app.NewServer(config, db)
	server.Start()
}
