package main

import (
	"crypto"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	HashAsm    = crypto.SHA512
	BCryptCost = bcrypt.DefaultCost
	SecretKey  = "SECRET_KEY" // TOOD: Must change this value
)

// Database
const (
	DSN      = "develop.db"
	LogLevel = 4
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DSN), &gorm.Config{})
	if err != nil {
		println("Failed to connect database. Reason: ", err.Error())
		panic("Failed to connect database")
	}
	fmt.Println("Connected to database")
	db.Logger.LogMode(LogLevel)
	return db
}
