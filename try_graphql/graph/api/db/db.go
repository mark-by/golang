package db

import (
	"database/sql"
	"fmt"
	"sync"
)

const (
	UserDb     = "test"
	PasswordDb = "123"
	NameDb     = "graph_test_db"
	SslMode     = "disable"
)

var once sync.Once

func Connect() (*sql.DB, error){
	var db *sql.DB
	var err error
	once.Do(func() {
		dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			UserDb, PasswordDb, NameDb, SslMode)
		db, _ = sql.Open("postgres", dbInfo)
		err = db.Ping()
	})
	return db, err
}

func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query)
	return db.Query(query, args...)
}

func MustExec(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}