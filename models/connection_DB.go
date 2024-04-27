package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to database
func Connection_DB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=quan1234 dbname=Server-Management-System port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
