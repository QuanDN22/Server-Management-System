package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to database
func ConnectionDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=quan1234 dbname=Server-Management-System port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("successfully connected database")
	}

	// delete table if it doesn't exist
	db.Migrator().DropTable(&Server{})
	db.Migrator().DropTable(&ServerDeleted{})

	// Auto migrate the Server model
	db.AutoMigrate(&Server{})
	db.AutoMigrate(&ServerDeleted{})

	return db, nil
}
