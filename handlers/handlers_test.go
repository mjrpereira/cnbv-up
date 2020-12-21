package handlers

import (
	"os"
	"testing"

	"github.com/mjrpereira/cnbv/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	server = Server{DB: db}

	os.Exit(m.Run())
}

func refreshUserTable() error {
	err := server.DB.Migrator().DropTable(&models.User{})
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	//log.Println("User table refreshed")
	return nil
}
