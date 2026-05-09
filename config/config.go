package config

import (
	"app/pkg/utils"
	"fmt"

	"github.com/spf13/viper"
)

type CommondArgs struct {
	ConfigDir  string
	ConfigFile string
}

type DbConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DB           string
	MaxOpenConns int
	MaxIdleConns int
	CharSet      string
	Collate      string
}

type RedisConfig struct {
	Host     string
	Port     int
	DB       int
	Password string
	Timeout  string
}

type AppConfig struct {
	Name string
	Port string
	Host string
}

type Config struct {
	CommondArgs CommondArgs
	DB          DbConfig
	Redis       RedisConfig
	App         AppConfig
}

// LoadConfigSource 加载配置文件
func LoadConfig(args *CommondArgs) (*Config, error) {
	config := &Config{}
	config.CommondArgs = *args

	if !utils.IsDirExists(args.ConfigDir) {
		return nil, fmt.Errorf("配置文件目录 %s 不存在", args.ConfigDir)
	}

	configPath := args.ConfigDir + "/" + args.ConfigFile
	if !utils.IsFileExists(configPath) {
		return nil, fmt.Errorf("配置文件 %s 不存在", configPath)
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件 %s 失败: %w", configPath, err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("解析配置文件 %s 失败: %w", configPath, err)
	}
	return config, nil
}
