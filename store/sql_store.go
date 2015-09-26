package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	l4g "github.com/macedo/category_service-go/Godeps/_workspace/src/code.google.com/p/log4go"
	"time"
)

type SqlStore struct {
	db *gorm.DB
}

func NewSqlStore() Store {

	sqlStore := &SqlStore{}
	sqlStore.db = getConnection()

	return sqlStore
}

func getConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://category_service-go:@127.0.0.1:5432/category_service-go?sslmode=disable")
	if err != nil {
		l4g.Critical("Failed to open sql connection to err:%v", err)
		time.Sleep(time.Second)
		panic("Failed to open sql connection " + err.Error())
	}

	db.DB()

	l4g.Info("Pinging sql postgres database")
	err = db.DB().Ping()
	if err != nil {
		l4g.Critical("Failed to ping db err:%v", err)
		time.Sleep(time.Second)
		panic("Failed  to open sql connection " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return &db
}

func (ss SqlStore) Close() {
	l4g.Info("Closing SqlStore")
	ss.db.Close()
}
