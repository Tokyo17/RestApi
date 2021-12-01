package products

import (
	"context"
	"time"
)

//MODEL DATA UNTUK BISNIS

type ProductDomain struct {
	ID             int
	Name           string
	Price          float64
	Picture_url    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Product_typeID int
	Product_type   Product_typeDomain
}

type Product_typeDomain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	FilterByType(ctx context.Context, typeid int) ([]ProductDomain, error)

	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)

	UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error)
	UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error)
}
type Repository interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	FilterByType(ctx context.Context, typeid int) ([]ProductDomain, error)

	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)

	UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error)
	UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error)
}
