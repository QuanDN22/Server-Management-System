package main

import (
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/data"
	"github.com/QuanDN22/Server-Management-System/models"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	// Connection to database
	db, err = models.Connection_DB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize database
	err = models.Init_DB(db)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// add_data into database
	data.AddData_Init(db)

	var servers []models.Server
	result := db.Find(&servers)
	fmt.Println(result.RowsAffected, result.Error)
}
