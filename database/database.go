package database

import (
	"github.com/danielsoro/amanda/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func Connect() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open("sqlite3", "./database/amanda.db")
		if err != nil {
			panic("failed to connect database")
		}
		instance = db
		db.AutoMigrate(&models.Phrase{})
	})

	return instance
}
