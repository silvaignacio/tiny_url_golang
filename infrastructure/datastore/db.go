package datastore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pet-system/domain/model"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=dev_loyalty dbname=dev_url port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.TinyUrl{})

	return db
}
