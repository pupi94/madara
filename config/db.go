package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", Env.DBUsername, Env.DBPassword, Env.DBHostname, Env.DBPort, Env.DBDatabase)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.(time.Second * 10)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(50)

	DB.ShowSQL(Env.DBShowSQL)

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
