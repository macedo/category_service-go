package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func PGSession() gorm.DB {
	db, err := gorm.Open("postgres", databaseURL)
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
