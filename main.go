package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/QuanDN22/Server-Management-System/models"
	"github.com/xuri/excelize/v2"
)

func main() {
	// Connection to database
	db := models.Connection_DB()

	// // Initialize database
	// models.Init_DB(db)

	// // add_data into database
	// models.AddData_Init(db)

	var servers []models.Server
	result := db.Find(&servers)
	fmt.Println(result.RowsAffected, result.Error)

	// export server into excel file
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Error exporting server file %v", err)
		}
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Fatalf("Error creating sheet file %v", err)
	}


	// Set value of a row
	err = f.SetSheetRow("Sheet1", "A1", &[]interface{}{
		"Server_ID",
		"Server_Name",
		"Server_IPv4",
		"Server_Status",
		"Server_CreatedAt",
		"Server_UpdatedAt",
	})
	if err != nil {
		log.Fatalf("Error setting value of a row %v", err)
	}

	i := 2
	for _, server := range servers {
		if i == 1000 {
			break
		}

		location := "A" + strconv.Itoa(i)

		err := f.SetSheetRow("Sheet1", location, &[]interface{}{
			server.ID,
			server.Server_Name,
			server.Server_IPv4,
			server.Server_Status,
			server.CreatedAt,
			server.UpdatedAt,
		})
		if err != nil {
			log.Fatalf("Error setting value of a row %v", err)
		}

		i++
	}

	// Set active sheet of the workbook.
    f.SetActiveSheet(index)
    // Save spreadsheet by the given path.
    if err := f.SaveAs("./data/data_export_example.xlsx"); err != nil {
		log.Fatalf("Error saving spreadsheet file %v", err)
	}
}
