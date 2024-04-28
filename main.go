package main

import (
	"fmt"
	"log"

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

	// // Initialize database
	// err = models.Init_DB(db)
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// // add_data into database
	// data.AddData_Init(db)

	// datas := []models.Server{
	// 	{Server_Name: "server#1", Server_IPv4: "192.168.1.1", Server_Status: true},
	// 	{Server_Name: "server#2", Server_IPv4: "192.168.1.2", Server_Status: true},
	// 	{Server_Name: "server#3", Server_IPv4: "192.168.1.3", Server_Status: true},
	// 	{Server_Name: "server#4", Server_IPv4: "192.168.1.4", Server_Status: true},
	// 	{Server_Name: "server#5", Server_IPv4: "192.168.1.5", Server_Status: true},
	// 	{Server_Name: "server#6", Server_IPv4: "192.168.1.6", Server_Status: true},
	// 	{Server_Name: "server#7", Server_IPv4: "192.168.1.7", Server_Status: true},
	// 	{Server_Name: "server#8", Server_IPv4: "192.168.1.8", Server_Status: true},
	// 	{Server_Name: "server#9", Server_IPv4: "192.168.1.9", Server_Status: true},
	// 	{Server_Name: "server#10", Server_IPv4: "192.168.1.10", Server_Status: true},
	// }

	// for _, data := range datas {
	// 	result := db.Create(&data) // pass pointer of data to Create
	// 	fmt.Println(result.RowsAffected, result.Error)
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// fmt.Print("Number servers: ")
	// var servers []models.Server
	// // result := db.Model(&models.Server{}).Find(&servers)
	// result := db.Find(&servers)
	// fmt.Println(result.RowsAffected, result.Error)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// for _, server := range servers {
	// 	fmt.Println(server)
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()

	// fmt.Println("Delete")
	// var i = 0
	// for _, server := range servers {
	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Println(server)
	// 	result = db.Delete(&server)
	// 	fmt.Println(result.RowsAffected, result.Error)

	// 	var server_deleted = &models.ServerDeleted{
	// 		Server_ID:   server.ID,
	// 		Server_Name: server.Server_Name,
	// 		Server_IPv4: server.Server_IPv4,
	// 		CreatedAt:   server.CreatedAt,
	// 		UpdatedAt:   server.UpdatedAt,
	// 		DeletedAt:   server.DeletedAt.Time,
	// 	}
	// 	result = db.Create(server_deleted)
	// 	fmt.Println(result.RowsAffected, result.Error)

	// 	result = db.Unscoped().Delete(&server)
	// 	fmt.Println(result.RowsAffected, result.Error)
	// }

	server_test := models.Server{
		Server_Name:   "server#0",
		Server_IPv4:   "1.1.1.1",
		Server_Status: true,
	}

	result := db.Create(&server_test)
	fmt.Println(result.RowsAffected, result.Error)

	var server_test2 models.Server
	result = db.Where("server_name = ?", "server#0").First(&server_test2)
	fmt.Println(result.RowsAffected, result.Error)

	result = db.Where("server_name = ?", "server#0").Delete(&server_test2)
	fmt.Println(result.RowsAffected, result.Error)

	var server_deleted = &models.ServerDeleted{
		Server_ID:   server_test2.ID,
		Server_Name: server_test2.Server_Name,
		Server_IPv4: server_test2.Server_IPv4,
		CreatedAt:   server_test2.CreatedAt,
		UpdatedAt:   server_test2.UpdatedAt,
		DeletedAt:   server_test2.DeletedAt.Time,
	}
	result = db.Create(server_deleted)
	fmt.Println(result.RowsAffected, result.Error)

	result = db.Where("server_name = ?", "server#0").Unscoped().Delete(&server_test)
	fmt.Println(result.RowsAffected, result.Error)
}
