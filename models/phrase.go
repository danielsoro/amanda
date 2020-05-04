package models

import (
	"github.com/jinzhu/gorm"
)

type Phrase struct {
	gorm.Model
	Value string `json:"phrase"`
}
