package database

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	once     sync.Once
	instance *gorm.DB
)

// Session return the gorm conn
func Session() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open("sqlite3", "./database/amanda.db")
		if err != nil {
			panic("failed to connect database")
		}
		instance = db
	})

	return instance
}
