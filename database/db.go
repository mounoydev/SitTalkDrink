package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const HOST = "localhost"
const PORT = "3307"
const USER_DB = "root"
const PASSWD_DB = "1234"
const DB_NAME = "testgodb"

var Db *gorm.DB
func Initdb() *gorm.DB {
	Db = dbConnect()
	return Db
}

func dbConnect() *gorm.DB {
	var err error
	dbUrl := USER_DB + ":" + PASSWD_DB + "@tcp" + "(" + HOST + ":" + PORT + ")/" + DB_NAME + "?parseTime=true&loc=Local"
	fmt.Println("dbURL: ", dbUrl)
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection error: %v", err)
	}

	return db
}
