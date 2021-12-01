package product

import (
	"context"
	"final/business/products"
	"gorm.io/gorm"
	"strings"
)

type MysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) products.Repository {
	return &MysqlProductRepository{
		Conn: conn,
	}
}

func (rep MysqlProductRepository) Get(ctx context.Context) ([]products.ProductDomain, error) {
	var product []Product

	result := rep.Conn.Preload("Product_type").Find(&product)

	if result.Error != nil {
		return nil, result.Error
	}
	return ToListDomain(product), nil
}

func (rep MysqlProductRepository) UploadType(ctx context.Context, domain products.Product_typeDomain) (products.Product_typeDomain, error) {
	var newProductType Product_type
	newProductType.Name = domain.Name
	result := rep.Conn.Create(&newProductType)
	if result.Error != nil {
		return products.Product_typeDomain{}, result.Error
	}
	return products.Product_typeDomain(newProductType), nil

}

func (rep MysqlProductRepository) FilterByType(ctx context.Context, typeid int) ([]products.ProductDomain, error) {
	var product []Product
	result := rep.Conn.Preload("Product_type").Where("product_type_id = ?", typeid).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToListDomain(product), nil
}

func (rep MysqlProductRepository) UploadProduct(ctx context.Context, productdomain products.ProductDomain) (products.ProductDomain, error) {
	var newProduct Product
	newProduct.Name = strings.ToLower(productdomain.Name)
	newProduct.Price = productdomain.Price
	newProduct.Picture_url = productdomain.Picture_url
	newProduct.Product_typeID = productdomain.Product_typeID
	result := rep.Conn.Create(&newProduct)
	if result.Error != nil {
		return products.ProductDomain{}, result.Error
	}
	return newProduct.ToDomain(), nil
}

func (rep MysqlProductRepository) UpdateProduct(ctx context.Context, domain products.ProductDomain, id int) (products.ProductDomain, error) {
	var product Product
	result := rep.Conn.First(&product, "id = ?", id)
	product.Name = domain.Name
	product.Price = domain.Price
	product.Picture_url = domain.Picture_url
	result.Save(&product)
	if result.Error != nil {
		return products.ProductDomain{}, result.Error
	}
	return product.ToDomain(), nil

}
