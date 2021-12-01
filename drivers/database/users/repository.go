package users

import (
	"context"
	"errors"
	"final/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep MysqlUserRepository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user User
	user.Name = domain.Name
	user.Email = domain.Email
	user.Password = domain.Password
	user.Phone_number = domain.Phone_number
	user.Address = domain.Address
	user.Picture_url = domain.Picture_url

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil

}

func (rep MysqlUserRepository) Details(ctx context.Context, id int) (users.Domain, error) {
	var user User
	result := rep.Conn.Find(&user, "id = ?", id)
	if result.Error != nil {
		return users.Domain{}, errors.New("invalid payment, please enter number same as total price")
	}
	return user.ToDomain(), nil
}

func (rep MysqlUserRepository) Login(ctx context.Context, email, password string) (users.Domain, error) {
	var user User

	result := rep.Conn.First(&user, "email = ? AND password = ?",
		email, password)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}
