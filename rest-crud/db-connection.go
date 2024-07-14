package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:root@tcp(127.0.0.1:3306)/employee?parseTime=true"
// var err error

func DataMigration() {
	Database, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("DB Connection Failed")
	}

	Database.AutoMigrate(&Employee{})
}