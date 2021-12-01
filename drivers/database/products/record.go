package product

import (
	"final/business/products"
	"time"
)

type Product struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"index"`
	Price          float64
	Picture_url    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Product_typeID int
	Product_type   Product_type
}

type Product_type struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (product *Product_type) ToDomain() products.Product_typeDomain {
	return products.Product_typeDomain{
		ID:        product.ID,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (product *Product) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		ID:             product.ID,
		Name:           product.Name,
		Price:          product.Price,
		Picture_url:    product.Picture_url,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		Product_typeID: product.Product_typeID,
		Product_type:   product.Product_type.ToDomain(),
	}
}

func ToListDomain(data []Product) (result []products.ProductDomain) {
	result = []products.ProductDomain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
