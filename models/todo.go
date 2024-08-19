package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID        uint64
	Item      string
	Completed int
	gorm.Model
}
