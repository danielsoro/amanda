package database

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	once     sync.Once
	instance *gorm.DB
)

// Db is the struct to provide the session of the data source
type Db struct {
}

// Session return the gorm conn
func (d Db) Session() *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open("sqlite3", "./database/amanda.db")
		if err != nil {
			panic("failed to connect database")
		}
		instance = db
	})

	return instance
}
