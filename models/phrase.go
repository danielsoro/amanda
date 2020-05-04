package models

import (
	"github.com/danielsoro/amanda/database"
	"github.com/jinzhu/gorm"
)

type Phrase struct {
	gorm.Model
	Value string `json:"phrase"`
}

func GetPhrases() []Phrase {
	var phrases []Phrase
	database.Connect().Find(&phrases)
	return phrases
}
