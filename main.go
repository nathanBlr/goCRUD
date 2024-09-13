package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type User struct {
	gorm.Model
	Name string
	Email string
    Password string
	ID int
}

func main(){
	db, err := gorm.Open(sqlite.Open(""))
}