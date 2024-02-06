package db

import (
	"log"

	"github.com/Doni-githu/todo-app/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	sqlDB, err2 := db.DB()
	if err2 != nil {
		return nil, err2
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	db.AutoMigrate(&models.Person{})

	log.Fatalln()

	return db, nil
}
