package transactions

import (
	"context"
	"time"
)

type Payment_MethodDomain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShipmentDomain struct {
	ID             int
	Name           string
	Shipment_Price float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Transaction_DetailDomain struct {
	UserID         int
	StatusShipment string
	TransactionID  int
	ProductID      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TransactionDomain struct {
	ID               int
	UserID           int
	Status           string
	ProductID        int
	Product          ProductDomain
	Payment_MethodID int
	Payment_Method   Payment_MethodDomain
	ShipmentID       int
	Shipment         ShipmentDomain
	Total_Price      float64
	Quantity         int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

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
	Add(ctx context.Context, domain TransactionDomain) (TransactionDomain, error)
	DetailSC(ctx context.Context, id int) ([]TransactionDomain, error)

	GetTransDetail(ctx context.Context, userid, transid int) (Transaction_DetailDomain, TransactionDomain, error)

	Pay(ctx context.Context, transactionid int, amount float64) (TransactionDomain, error)

	AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error)
	GetPM(ctx context.Context) ([]Payment_MethodDomain, error)

	AddShipment(ctx context.Context, domain ShipmentDomain) (ShipmentDomain, error)
	GetShipment(ctx context.Context) ([]ShipmentDomain, error)
}

type Repository interface {
	Add(ctx context.Context, domain TransactionDomain) (TransactionDomain, error)
	DetailSC(ctx context.Context, id int) ([]TransactionDomain, error)

	GetTransDetail(ctx context.Context, userid, transid int) (Transaction_DetailDomain, TransactionDomain, error)

	Pay(ctx context.Context, transactionid int, amount float64) (TransactionDomain, error)

	AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error)
	GetPM(ctx context.Context) ([]Payment_MethodDomain, error)

	AddShipment(ctx context.Context, domain ShipmentDomain) (ShipmentDomain, error)
	GetShipment(ctx context.Context) ([]ShipmentDomain, error)
}
