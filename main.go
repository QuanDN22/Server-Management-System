package main

import (
	"fmt"

	"github.com/QuanDN22/Server-Management-System/models"
)

func main() {
	// Connection to database
	db := models.Connection_DB()

	// Initialize database
	models.Init_DB(db)

	// add_data into database
	models.AddData_Init(db)

	var servers []models.Server
	result := db.Find(&servers)
	fmt.Println(result.RowsAffected, result.Error)
}
