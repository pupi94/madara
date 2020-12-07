package config

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", Env.DBUsername, Env.DBPassword, Env.DBHostname, Env.DBPort, Env.DBDatabase)

	var err error
	DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Second * 10)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(50)

	DB.ShowSQL(Env.DBShowSQL)

	//log := logrus.New()
	//level, err := logrus.ParseLevel(Env.LogLevel)
	//if err != nil {
	//	panic(err)
	//}
	//log.SetLevel(level)
	//DB.SetLogger(log)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
