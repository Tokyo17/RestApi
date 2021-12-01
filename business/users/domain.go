package users

import (
	"context"
	"time"
)

type Domain struct {
	ID           int
	Name         string
	Email        string
	Password     string
	Token        string
	Phone_number int
	Address      string
	Picture_url  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type UseCase interface {
	Details(ctx context.Context, id int) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
}

type Repository interface {
	Details(ctx context.Context, id int) (Domain, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
}
