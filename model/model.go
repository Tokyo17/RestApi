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

type Products struct {
	ID        int
	TypeId    int `gorm:"ForeignKey:ID"`
	Type      Type
	Name      string
	Foto      string
	Size      string
	Color     string
	Deleted   gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Type struct {
	ID        int
	Name      string
	Deleted   gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
