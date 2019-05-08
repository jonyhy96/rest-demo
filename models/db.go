package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var conn *gorm.DB

// Connect connect to the postgresql
func Connect(host string, database string, user string, pass string) (db *gorm.DB, err error) {
	dbConnString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, pass)
	db, err = gorm.Open("postgres", dbConnString)
	db.DB().SetMaxIdleConns(5)
	conn = db
	return
}

// GetDb return the db instance
func GetDb() *gorm.DB {
	return conn
}
