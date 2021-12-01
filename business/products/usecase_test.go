package products_test

import (
	"context"
	"final/business/products"
	"final/business/products/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository mocks.Repository
var productService products.UseCase
var productDomain products.ProductDomain
var productTypeDomain products.Product_typeDomain

func setup() {
	productService = products.NewProductUseCase(&productRepository, time.Hour*1)
	productDomain = products.ProductDomain{
		ID:             1,
		Name:           "Gamis Biru",
		Price:          50000,
		Picture_url:    "www.google.com",
		Product_typeID: 1,
		Product_type:   productTypeDomain,
	}

	productTypeDomain = products.Product_typeDomain{
		ID:   1,
		Name: "Gamis",
	}
}

func TestGet(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Get", func(t *testing.T) {
		productRepository.On("Get",
			mock.Anything).Return([]products.ProductDomain{}, nil).Once()
		_, err := productService.Get(context.Background())
		assert.Nil(t, err)

	})
}

func TestUploadType(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UploadType", func(t *testing.T) {
		productRepository.On("UploadType",
			mock.Anything,
			mock.Anything).Return(products.Product_typeDomain{}, nil).Once()
		_, err := productService.UploadType(context.Background(), products.Product_typeDomain{
			ID:   1,
			Name: "Gamis",
		})
		assert.Nil(t, err)

	})

}

func TestUpdateProduct(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UpdateProduct", func(t *testing.T) {
		productRepository.On("UpdateProduct",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UpdateProduct(context.Background(), products.ProductDomain{
			ID:             1,
			Name:           "Gamis Biru",
			Price:          50000,
			Picture_url:    "www.google.com",
			Product_typeID: 1,
			Product_type:   productTypeDomain,
		}, 1)
		assert.Nil(t, err)

	})

}

func TestUploadProduct(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid UploadProduct", func(t *testing.T) {
		productRepository.On("UploadProduct",
			mock.Anything,
			mock.Anything).Return(products.ProductDomain{}, nil).Once()
		_, err := productService.UploadProduct(context.Background(), products.ProductDomain{
			ID:             1,
			Name:           "Gamis Biru",
			Price:          50000,
			Picture_url:    "www.google.com",
			Product_typeID: 1,
			Product_type:   productTypeDomain,
		})
		assert.Nil(t, err)

	})

}
