package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func PostgresSession() gorm.DB {
	db, err := gorm.Open("postgres", "user=category_service-go dbname=category_service-go sslmode=disable")

	if err != nil {
		log.Fatal("Create session: %s\n", err.Error())
	}

	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}
