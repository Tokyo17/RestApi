package products

import (
	"context"
	"errors"
	"time"
)

type ProductUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewProductUseCase(repo Repository, timeOut time.Duration) UseCase {
	return &ProductUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
	}
}

//Mengisi
func (uc *ProductUseCase) Get(ctx context.Context) ([]ProductDomain, error) {
	product, err := uc.Repo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

//Func untuk memanggi
func (uc *ProductUseCase) UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error) {
	if domain.Name == "" {
		return Product_typeDomain{}, errors.New("product type name is empty")
	}
	productType, err := uc.Repo.UploadType(ctx, domain)
	if err != nil {
		return Product_typeDomain{}, err
	}
	return productType, nil
}

func (uc *ProductUseCase) FilterByType(ctx context.Context, typeid int) ([]ProductDomain, error) {
	if typeid == 0 {
		return nil, errors.New("typeid is empty")
	}
	product, err := uc.Repo.FilterByType(ctx, typeid)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (uc *ProductUseCase) UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error) {
	if productdomain.Name == "" {
		return ProductDomain{}, errors.New("product name is empty")
	}
	if productdomain.Price == 0 {
		return ProductDomain{}, errors.New("price is empty")
	}
	if productdomain.Price < 0 {
		return ProductDomain{}, errors.New("invalid price")
	}
	if productdomain.Picture_url == "" {
		return ProductDomain{}, errors.New("picture is empty")
	}
	if productdomain.Product_typeID == 0 {
		return ProductDomain{}, errors.New("product type id is empty")
	}
	product, err := uc.Repo.UploadProduct(ctx, productdomain)
	if err != nil {
		return ProductDomain{}, err
	}
	return product, nil
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error) {
	if domain.Price < 0 {
		return ProductDomain{}, errors.New("invalid price")
	}
	updateProduct, err := uc.Repo.UpdateProduct(ctx, domain, id)
	if err != nil {
		return ProductDomain{}, nil
	}
	return updateProduct, nil
}
