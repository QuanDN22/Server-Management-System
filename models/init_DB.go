package models

import (
	"log"

	"gorm.io/gorm"
)

func Init_DB(db *gorm.DB) {
	var err error

	// delete table if it doesn't exist
	err = db.Migrator().DropTable(&Server{})
	if err != nil {
		log.Fatalf("Failed to drop table servers: %v", err)
	}

	err = db.Migrator().DropTable(&ServerDeleted{})
	if err != nil {
		log.Fatalf("Failed to drop table server_deleteds: %v", err)
	}

	// Auto migrate the Server model
	err = db.AutoMigrate(&Server{})
	if err != nil {
		log.Fatalf("Failed to migrate servers datable: %v", err)
	}

	err = db.AutoMigrate(&ServerDeleted{})
	if err != nil {
		log.Fatalf("Failed to migrate server_deleteds datable: %v", err)
	}
}
