package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uint64
	Item      string
	Completed int
}
