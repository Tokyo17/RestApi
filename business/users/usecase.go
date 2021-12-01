package users

import (
	"context"
	"errors"
	"final/app/middleware"
	"time"
)

type UserUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JwtToken       middleware.ConfigJWT
}

func NewUserUseCase(repo Repository, timeOut time.Duration, token middleware.ConfigJWT) UseCase {
	return &UserUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
		JwtToken:       token,
	}
}

func (uc *UserUseCase) Details(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.Details(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) Register(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Name == "" {
		return Domain{}, errors.New("name is empty")
	}
	if domain.Phone_number == 0 {
		return Domain{}, errors.New("phone number error awali dengan 62")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email is empty")
	}

	if domain.Address == "" {
		return Domain{}, errors.New("Address is empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("Password is empty")
	}
	if domain.Picture_url == "" {
		return Domain{}, errors.New("Picture_url is empty")
	}

	user, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email is empty")
	}

	if password == "" {
		return Domain{}, errors.New("password is empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)
	var fail error
	user.Token, fail = uc.JwtToken.GenerateToken(user.ID)
	if fail != nil {
		return Domain{}, fail
	}
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
