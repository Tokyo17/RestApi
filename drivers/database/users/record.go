package users

import (
	"final/business/users"
	"time"
)

type User struct {
	ID int `gorm:"primaryKey"`

	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Phone_number int
	Address      string
	Picture_url  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Phone_number: user.Phone_number,
		Address:      user.Address,
		Picture_url:  user.Picture_url,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:           domain.ID,
		Name:         domain.Name,
		Email:        domain.Email,
		Password:     domain.Password,
		Phone_number: domain.Phone_number,
		Address:      domain.Address,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
