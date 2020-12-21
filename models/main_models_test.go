package models

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

var server Server

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	server.DB = db
	os.Exit(m.Run())
}

func refreshUserTable() error {
	err := server.DB.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	//log.Println("User table refreshed")
	return nil
}
