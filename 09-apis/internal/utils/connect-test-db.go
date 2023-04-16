package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToTestDatabase(models ...interface{}) (tx *gorm.DB, db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	tx = db.Begin()
	tx.AutoMigrate(models...)

	return tx, db
}
