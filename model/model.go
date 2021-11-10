package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int
	Name      string `JSON:"name"`
	Email     string
	Address   string
	Password  string
	Deleted   gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}