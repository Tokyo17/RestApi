package users

import (
	"final/business/transactions"
	"time"
)

type User struct {
	ID int `gorm:"primaryKey" json:"id"`

	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Phone_number int
	Picture_url  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Product_type struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
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
type Transaction struct {
	ID               int `gorm:"primaryKey"`
	UserID           int
	Status           string
	ProductID        int
	Product          Product
	Payment_MethodID int
	Payment_Method   Payment_Method
	ShipmentID       int
	Shipment         Shipment
	Quantity         int
	Total_Price      float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Payment_Method struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Shipment struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Shipment_Price float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Transaction_Detail struct {
	UserID         int
	StatusShipment string
	TransactionID  int `gorm:"primaryKey"`
	Transaction    Transaction
	ProductID      int `gorm:"primaryKey"`
	Product        Product
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (transaction_details *Transaction_Detail) ToDomain() transactions.Transaction_DetailDomain {
	return transactions.Transaction_DetailDomain{

		UserID:         transaction_details.UserID,
		StatusShipment: transaction_details.StatusShipment,
		TransactionID:  transaction_details.TransactionID,
		ProductID:      transaction_details.ProductID,
		CreatedAt:      transaction_details.CreatedAt,
		UpdatedAt:      transaction_details.UpdatedAt,
	}
}

func (Transaction *Transaction) ToDomain() transactions.TransactionDomain {
	return transactions.TransactionDomain{
		ID:               Transaction.ID,
		UserID:           Transaction.UserID,
		Status:           Transaction.Status,
		Payment_MethodID: Transaction.Payment_MethodID,
		Payment_Method:   Transaction.Payment_Method.ToDomain(),
		ProductID:        Transaction.ProductID,
		Product:          Transaction.Product.ToDomain(),
		ShipmentID:       Transaction.ShipmentID,
		Shipment:         Transaction.Shipment.ToDomain(),
		Quantity:         Transaction.Quantity,
		Total_Price:      Transaction.Total_Price,
		CreatedAt:        Transaction.CreatedAt,
		UpdatedAt:        Transaction.UpdatedAt,
	}
}

func (payment_method *Payment_Method) ToDomain() transactions.Payment_MethodDomain {
	return transactions.Payment_MethodDomain{
		ID:        payment_method.ID,
		Name:      payment_method.Name,
		CreatedAt: payment_method.CreatedAt,
		UpdatedAt: payment_method.UpdatedAt,
	}
}
func (shipment *Shipment) ToDomain() transactions.ShipmentDomain {
	return transactions.ShipmentDomain{
		ID:             shipment.ID,
		Name:           shipment.Name,
		Shipment_Price: shipment.Shipment_Price,
		UpdatedAt:      shipment.UpdatedAt,
		CreatedAt:      shipment.CreatedAt,
	}
}

func ListSCToDomain(data []Transaction) (result []transactions.TransactionDomain) {
	for _, SC := range data {
		result = append(result, SC.ToDomain())
	}
	return
}

func ListShipmentToDomain(data []Shipment) (result []transactions.ShipmentDomain) {
	for _, Shipment := range data {
		result = append(result, Shipment.ToDomain())
	}
	return
}

func ListPMToDomain(data []Payment_Method) (result []transactions.Payment_MethodDomain) {
	for _, PM := range data {
		result = append(result, PM.ToDomain())
	}
	return
}

func (product *Product) ToDomain() transactions.ProductDomain {
	return transactions.ProductDomain{
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
func (product *Product_type) ToDomain() transactions.Product_typeDomain {
	return transactions.Product_typeDomain{
		ID:        product.ID,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
