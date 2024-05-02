package main

import (
	"github.com/QuanDN22/Server-Management-System/logger"
	"github.com/QuanDN22/Server-Management-System/models"
)

func main() {
	// run log
	logFile := logger.LogFile()
	defer logFile.Close()

	// Connection to database
	db := models.Connection_DB()

	// Initialize database
	models.Init_DB(db)

	// add_data into database
	models.AddData_Init(db)

	// export data from database into excel file
	models.ExportData_Example(db)
}
