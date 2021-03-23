package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", Env.DBUsername, Env.DBPassword, Env.DBHostname, Env.DBPort, Env.DBDatabase)

	var err error
	DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	// 设置数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(50)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
}
