package database

import (
	"fmt"
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
			panic(fmt.Errorf("failed to connect database: %v", err.Error()))
		}
		instance = db
	})

	return instance
}
