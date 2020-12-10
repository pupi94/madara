package main

import (
	"github.com/pupi94/madara/config"
)

type SchemaMigration struct {
	Version string `xorm:"version"`
}

func main() {
	config.InitDB()

	s := &SchemaMigration{Version: "20201207105022"}
	_, err := config.DB.Insert(s)
	if err != nil {
		panic(err)
	}
}
