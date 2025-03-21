package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	if dbConn == nil {
		log.Panicf("Failed to connect DB :%v", dbConn)
	}
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Panicf("Failed to migrate db: %v", err)
	}
	fmt.Println("Successfully Migrated")
}
