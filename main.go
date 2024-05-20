package main

import (
	"log"
)

func main() {
	// run log
	// logFile := logger.LogFile()
	// defer logFile.Close()

	// // Connection to database
	// db := models.Connection_DB()

	// // Initialize database
	// models.Init_DB(db)

	// // add_data into database
	// models.AddData_Init(db)

	// // export data from database into excel file
	// models.ExportData_Example(db)

	// c, err := config.NewConfig("./pkg/config", ".env.auth")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Println("config parsed...")
	// fmt.Println(c)
	// fmt.Println(c.ServiceName)
	// fmt.Println(c.GrpcAddr)
	// fmt.Println(c.GrpcPort)
}
