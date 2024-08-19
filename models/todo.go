package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint64         `json:"id"`
	Item      string         `json:"item"`
	Completed int            `json:"completed"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
