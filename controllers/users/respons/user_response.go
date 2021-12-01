package respons

import (
	"final/business/users"
	"time"
)

type UserResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	Phone_number int       `json:"phone_number"`
	Address      string    `json:"Address"`
	Picture_url  string    `json:"picture_url"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:           domain.ID,
		Name:         domain.Name,
		Email:        domain.Email,
		Token:        domain.Token,
		Phone_number: domain.Phone_number,
		Address:      domain.Address,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type UserDetailResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone_number int       `json:"phone_number"`
	Address      string    `json:"Address"`
	Picture_url  string    `json:"picture_url"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func FromDetailDomain(domain users.Domain) UserDetailResponse {
	return UserDetailResponse{
		ID:           domain.ID,
		Name:         domain.Name,
		Email:        domain.Email,
		Phone_number: domain.Phone_number,
		Address:      domain.Address,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
