package admins_test

import (
	"context"
	"final/app/middleware"
	"final/business/admins"
	"final/business/admins/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var adminRepository mocks.Repository
var adminSevice admins.UseCase
var adminDomain admins.AdminDomain
var adminToken middleware.ConfigJWT

func setup() {
	adminSevice = admins.NewAdminUseCase(&adminRepository, time.Hour*1, adminToken)
	adminDomain = admins.AdminDomain{
		ID:       1,
		Name:     "Yuli",
		Email:    "yuli@gmail.com",
		Password: "yuli123",
		Token:    "123",
	}
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(adminDomain, nil).Once()

		admin, err := adminSevice.Login(context.Background(), "yuli@gmail.com", "yuli123")

		assert.Nil(t, err)
		assert.Equal(t, "Yuli", admin.Name)
	})

}

func TestRegister(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		adminRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(adminDomain, nil).Once()

		_, err := adminSevice.Register(context.Background(), admins.AdminDomain{
			Name:     "Yuli",
			Email:    "yuli@gmail.com",
			Password: "yuli123",
		})

		assert.Nil(t, err)
	})
}
