package cmd

import (
	"context"
	"errors"
	"github.com/pupi94/madara/db"
)

func DbMigrate(ctx context.Context, direction string, step int) error {
	if direction == "" || direction == "up" {
		db.MigrateUp()
	} else if direction == "down" {
		db.MigrateDown(step)
	} else {
		return errors.New("database migrate args error")
	}
	return nil
}

func GenerateMigration(ctx context.Context, name string) error {
	db.GenerateMigration(name)
	return nil
}
