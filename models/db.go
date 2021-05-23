package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(db_url string){
  DSN := db_url

  db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err !=nil{
		panic("Failed to connect to database")
	}
  db.AutoMigrate(&User{}, &UserAccount{}, &Transactions{})
	DB = db
}