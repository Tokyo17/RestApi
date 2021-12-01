package respons

import (
	"final/business/products"
	"strings"
	"time"
)

// MODEL UNTUK respons

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Picture_url string    `json:"picture_url"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	// Product_typeID int                  `json:"product_typeid"`
	Product_type Product_typeResponse `json:"product_type"`
}

type Product_typeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type UploadProductResponse struct {
	ID int `json:"id"`

	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	Picture_url    string    `json:"picture_url"`
	Product_typeID int       `json:"product_typeid"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// MENGISI STRUCT ProductResponse
func FromDomain(domain products.ProductDomain) ProductResponse {
	return ProductResponse{
		ID:          domain.ID,
		Name:        strings.Title(domain.Name),
		Price:       domain.Price,
		Picture_url: domain.Picture_url,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		// Product_typeID: domain.Product_typeID,
		Product_type: TypeFromDomain(domain.Product_type),
	}
}

func TypeFromDomain(domain products.Product_typeDomain) Product_typeResponse {
	return Product_typeResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func ProductFromDomain(domain products.ProductDomain) UploadProductResponse {
	return UploadProductResponse{
		ID:             domain.ID,
		Name:           strings.Title(domain.Name),
		Price:          domain.Price,
		Picture_url:    domain.Picture_url,
		Product_typeID: domain.Product_typeID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

//MELOOP DATA DAN MEMASUKKAN KE FromDomain
func ListFromDomain(data []products.ProductDomain) (result []ProductResponse) {
	result = []ProductResponse{}
	for _, products := range data {
		products.Name = strings.Title(products.Name)
		result = append(result, FromDomain(products))
	}
	return
}
