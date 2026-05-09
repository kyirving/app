package database

import (
	"app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(config *config.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB, config.CharSet, config.Collate)

	dnConfig := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(dsn), dnConfig)
	if err != nil {
		return nil, err
	}

	DB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池大小
	DB.SetMaxOpenConns(config.MaxOpenConns)
	DB.SetMaxIdleConns(config.MaxIdleConns)
	return db, nil
}
