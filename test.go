package main

import (
	"fmt"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/model"
)

func main() {
	config.InitDB()

	p := model.Product{Title: "test", Description: "212", StoreID: 1212}

	fmt.Println(p.ID)

	db := config.DB
	res := db.Create(&p)
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println(p.ID)
}
