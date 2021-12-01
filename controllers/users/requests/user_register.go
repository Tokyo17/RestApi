package requests

import (
	"final/business/users"
	// "time"
)

type UserRegister struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone_number int    `json:"phone_number"`
	Address      string `json:"address"`
	Picture_url  string `json:"picture_url"`
}

func (user *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Phone_number: user.Phone_number,
		Address:      user.Address,
		Picture_url:  user.Picture_url,
	}
}
