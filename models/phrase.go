package models

import (
	"github.com/danielsoro/amanda/database"
	"github.com/jinzhu/gorm"
)

// Phrase struct for show prhases on the page
type Phrase struct {
	gorm.Model
	Value string `json:"phrase"`
}

// GetPhrases returns all phrases
func (p Phrase) GetPhrases() []Phrase {
	var phrases []Phrase
	database.Db{}.Session().Find(&phrases)
	return phrases
}
