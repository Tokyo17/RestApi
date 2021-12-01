package users

import (
	"context"
	"errors"
	"final/business/transactions"
	"fmt"

	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) transactions.Repository {
	return &MysqlTransactionRepository{
		Conn: conn,
	}
}

func (rep MysqlTransactionRepository) Add(ctx context.Context, domain transactions.TransactionDomain) (transactions.TransactionDomain, error) {
	var transaction Transaction
	var shipment Shipment
	var payment_method Payment_Method
	var product Product
	productSearch := rep.Conn.Where("id = ?", domain.ProductID).First(&product)
	if productSearch.Error != nil {
		return transactions.TransactionDomain{}, productSearch.Error
	}

	shipmentChoose := rep.Conn.First(&shipment, "id = ?", domain.ShipmentID)
	if shipmentChoose.Error != nil {
		return transactions.TransactionDomain{}, shipmentChoose.Error
	}
	transaction.Status = "Belum Dibayar"
	transaction.UserID = domain.UserID
	transaction.ProductID = domain.ProductID
	transaction.Payment_MethodID = domain.Payment_MethodID
	transaction.ShipmentID = domain.ShipmentID
	transaction.Quantity = domain.Quantity
	transaction.Total_Price = float64(domain.Quantity)*product.Price + shipment.Shipment_Price

	result := rep.Conn.Preload("Product").Preload("Payment_Method").Preload("Shipment").Create(&transaction)
	if result.Error != nil {
		return transactions.TransactionDomain{}, result.Error
	}
	rep.Conn.Where("id = ?", domain.ShipmentID).First(&shipment)
	rep.Conn.Where("id = ?", domain.Payment_MethodID).First(&payment_method)

	return transaction.ToDomain(), nil

}

func (rep MysqlTransactionRepository) DetailSC(ctx context.Context, id int) ([]transactions.TransactionDomain, error) {
	var listSC []Transaction
	result := rep.Conn.Preload("Product").Preload("Payment_Method").Preload("Shipment").Find(&listSC, "user_id = ?", id)
	if result.Error != nil {
		return []transactions.TransactionDomain{}, result.Error
	}
	return ListSCToDomain(listSC), nil
}

func (rep MysqlTransactionRepository) AddPM(ctx context.Context, domain transactions.Payment_MethodDomain) (transactions.Payment_MethodDomain, error) {
	var payment_method Payment_Method
	payment_method.Name = domain.Name
	result := rep.Conn.Create(&payment_method)
	if result.Error != nil {
		return transactions.Payment_MethodDomain{}, result.Error
	}
	return payment_method.ToDomain(), nil
}

func (rep MysqlTransactionRepository) GetPM(ctx context.Context) ([]transactions.Payment_MethodDomain, error) {
	var listPayment []Payment_Method
	result := rep.Conn.Find(&listPayment)
	if result.Error != nil {
		return []transactions.Payment_MethodDomain{}, result.Error
	}
	return ListPMToDomain(listPayment), nil
}

func (rep MysqlTransactionRepository) AddShipment(ctx context.Context, domain transactions.ShipmentDomain) (transactions.ShipmentDomain, error) {
	var shipment Shipment
	shipment.Name = domain.Name
	shipment.Shipment_Price = domain.Shipment_Price
	result := rep.Conn.Create(&shipment)
	if result.Error != nil {
		return transactions.ShipmentDomain{}, result.Error
	}
	return shipment.ToDomain(), nil
}

func (rep MysqlTransactionRepository) GetShipment(ctx context.Context) ([]transactions.ShipmentDomain, error) {
	var listshipment []Shipment
	result := rep.Conn.Find(&listshipment)
	if result.Error != nil {
		return []transactions.ShipmentDomain{}, result.Error
	}
	return ListShipmentToDomain(listshipment), nil
}

func (rep MysqlTransactionRepository) Pay(ctx context.Context, transactionid int, amount float64) (transactions.TransactionDomain, error) {
	var transaction Transaction

	chooseTransaction := rep.Conn.First(&transaction, "id = ?", transactionid)
	if chooseTransaction.Error != nil {
		return transactions.TransactionDomain{}, chooseTransaction.Error
	}

	if transaction.Total_Price == amount {

		updateStatus := rep.Conn.Preload("Shipment").Preload("Product").Preload("Payment_Method").First(&transaction, "id = ?", transactionid).Table("transactions").Where("id = ?", transactionid).Updates(map[string]interface{}{"status": "Dibayar"})
		if updateStatus.Error != nil {
			return transactions.TransactionDomain{}, updateStatus.Error
		}

		var detail Transaction_Detail
		detail.UserID = transaction.UserID
		detail.StatusShipment = "Dikirim"
		detail.TransactionID = transactionid
		detail.ProductID = transaction.ProductID
		createDetail := rep.Conn.Create(&detail)
		if createDetail.Error != nil {
			return transactions.TransactionDomain{}, createDetail.Error
		}

	} else {
		return transactions.TransactionDomain{}, errors.New("invalid payment, please enter number same as total price")
	}
	return transaction.ToDomain(), nil
}

func (rep MysqlTransactionRepository) GetTransDetail(ctx context.Context, userid, transid int) (transactions.Transaction_DetailDomain, transactions.TransactionDomain, error) {
	var detail Transaction_Detail
	var transaction Transaction

	fmt.Println(userid, transid)
	searchDetail := rep.Conn.Where("transaction_id = ? ", transid).Where("user_id = ? ", userid).Find(&detail)

	if searchDetail.Error != nil {
		return transactions.Transaction_DetailDomain{}, transactions.TransactionDomain{}, searchDetail.Error
	}
	searchTransaction := rep.Conn.Preload("Product").Preload("Shipment").Preload("Payment_Method").Find(&transaction, "id = ?", transid)
	if searchTransaction.Error != nil {
		return transactions.Transaction_DetailDomain{}, transactions.TransactionDomain{}, searchTransaction.Error
	}

	return detail.ToDomain(), transaction.ToDomain(), nil
}

// control.use.db.
