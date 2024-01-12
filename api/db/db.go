package db

import (
	"github.com/Shivd131/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// initialize the db connection
func InitDB() error {
	var err error
	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=myrestapi sslmode=disable password=secret")
	if err != nil {
		return err
	}

	// create tables based on the Item struct
	db.AutoMigrate(&models.Item{})

	return nil
}

// returns the dbinstance
func GetDB() *gorm.DB {
	return db
}
