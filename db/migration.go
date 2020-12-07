package db

import (
	"fmt"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/models"
)

func Migrate() {
	var err error
	db := config.DB
	fmt.Println(db.IsTableExist(models.Product{}))
	_, err = db.ImportFile("./db/migrations/20201207105022_create_products.sql")
	if err != nil {
		panic(err)
	}
}
