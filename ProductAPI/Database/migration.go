package Database

import (
	"log"
	"product-api/Models"
)

func MigrateDatabase() {

	err := DB.AutoMigrate(&Models.Product{})
	if err != nil {

		log.Println("Migration Faild!", err)
	}
	log.Println("Migration Success...")
}