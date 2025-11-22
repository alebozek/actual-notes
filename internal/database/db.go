package database

import (
	"log"
	"os"

	"github.com/alebozek/actual-notes/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBLogger = log.New(os.Stdout, "[DB] ", log.LstdFlags)

var DB *gorm.DB

func InitDB() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "notes.db" // set default db path
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		DBLogger.Fatal("failed to connect database: " + err.Error())
	}
	err = db.AutoMigrate(&models.User{}, &models.Session{}, &models.Note{}, &models.Folder{})
	if err != nil {
		DBLogger.Fatal("failed to migrate database: " + err.Error())
	}
	DB = db
	DBLogger.Println("database connected successfully")
}
