package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/models"
)

func main() {
	config.InitDB()

	p := models.Product{Title: "test", Description: "212", StoreID: uuid.New()}

	fmt.Println(p.ID)

	db := config.DB
	res := db.Create(&p)
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println(p.ID)
}
