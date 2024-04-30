package main

import (
	"github.com/QuanDN22/Server-Management-System/models"
)

func main() {
	// Connection to database
	db := models.Connection_DB()

	// Initialize database
	models.Init_DB(db)

	// add_data into database
	models.AddData_Init(db)

	// export data from database into excel file
	models.ExportData_Example(db)
}
