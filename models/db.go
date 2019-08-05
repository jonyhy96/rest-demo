package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	// import dialects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var conn *gorm.DB

// InitDB init DB
func InitDB() (db *gorm.DB, e error) {
	dbHost := beego.AppConfig.String("dbHost")
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName")
	dbPass := beego.AppConfig.String("dbPass")
	conn, err := Connect(dbHost, dbName, dbUser, dbPass)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

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
	var err error
	if conn == nil || conn.DB().Ping() != nil {
		conn, err = InitDB()
		if err != nil {
			glog.Error("db init error")
		}
	}
	return conn
}
