package models

import "gorm.io/gorm"

func Init_DB(db *gorm.DB) error {
	var err error

	// delete table if it doesn't exist
	err = db.Migrator().DropTable(&Server{})
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&ServerDeleted{})
	if err != nil {
		return err
	}

	// Auto migrate the Server model
	err = db.AutoMigrate(&Server{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&ServerDeleted{})
	if err != nil {
		return err
	}

	return nil
}
