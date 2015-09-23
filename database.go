package main

import (
	"log"

	"github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/jinzhu/gorm"
	_ "github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/lib/pq"
)

func DB() gorm.DB {
	db, err := gorm.Open("postgres", cfg.DatabaseURL)
	defer db.Close()

	if err != nil {
		log.Fatal("Create session: %s\n", err.Error())
	}

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}
