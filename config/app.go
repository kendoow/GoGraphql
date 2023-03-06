package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


var (
	db * gorm.DB
)


func Connect(){
	dbURL := "host=localhost user=postgres password=passwordzx dbname=test port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})	
	if err != nil{ 
		panic(err)
	}
	db = d
}

func GetBD() *gorm.DB  {
	return db
}