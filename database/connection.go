package database

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // nolint: gotype
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

/*
Connect connect to database
*/
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "triviago.db")
	// db, err := gorm.Open("postgres", "host=localhost dbname=trivia user=postgres password=b3g00dx sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
