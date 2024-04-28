package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to database
func Connection_DB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var varEnv map[string]string
	varEnv, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	var (
		host     = varEnv["PG_DATABASE_HOST"]
		user     = varEnv["PG_DATABASE_USER"]
		password = varEnv["PG_DATABASE_PASSWORD"]
		dbname   = varEnv["PG_DATABASE_DBNAME"]
		port     = varEnv["PG_DATABASE_PORT"]
	)

	portInt, err := strconv.Atoi(port)
	if err != nil {
		// Handle conversion error (e.g., invalid port format)
		panic(err)
	}

	// dsn :=  "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, portInt)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
