package models

import (
	"github.com/jinzhu/gorm"
)

// Phrase struct for show prhases on the page
type Phrase struct {
	gorm.Model
	Value string `json:"phrase"`
}
