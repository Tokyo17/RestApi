package users_test

import (
	"context"
	"final/app/middleware"
	"final/business/users"
	"final/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository mocks.Repository
var userService users.UseCase
var userDomain users.Domain
var userToken middleware.ConfigJWT

func setup() {
	userService = users.NewUserUseCase(&userRepository, time.Hour*1, userToken)
	userDomain = users.Domain{
		ID:           1,
		Name:         "Yuli",
		Email:        "yuli@gmail.com",
		Password:     "123",
		Token:        "123",
		Phone_number: 85641441299,
		Address:      "Bojonegoro",
		Picture_url:  "www.google.com",
	}

}

func TestRegister(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Yuli",
			Email:        "yuli@gmail.com",
			Password:     "123",
			Phone_number: 85641441299,
			Address:      "Bojonegoro",
			Picture_url:  "www.google.com",
		})
		assert.Nil(t, err)

	})
}

func TestDetails(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Details", func(t *testing.T) {
		userRepository.On("Details",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		user, err := userService.Details(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, "Yuli", user.Name)
	})
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		user, err := userService.Login(context.Background(), "yuli@gmail.com", "123")
		assert.Nil(t, err)
		assert.Equal(t, "Yuli", user.Name)
	})

}
