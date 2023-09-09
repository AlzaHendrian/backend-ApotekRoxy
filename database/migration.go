package database

import (
	"backend/models"
	"backend/pkg/postgresql"
	"fmt"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.Barang{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration success")
}
