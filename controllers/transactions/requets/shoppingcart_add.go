package requests

import "final/business/transactions"

type Shopping_CartAdd struct {
	UserID           int `json:"userid"`
	ProductID        int `json:"productid"`
	Payment_MethodID int `json:"payment_methodid"`
	ShipmentID       int `json:"shipmentid"`
	Quantity         int `json:"quantity"`
}

func (shopping_cartadd *Shopping_CartAdd) ToDomain() transactions.TransactionDomain {
	return transactions.TransactionDomain{
		UserID:           shopping_cartadd.UserID,
		ProductID:        shopping_cartadd.ProductID,
		Payment_MethodID: shopping_cartadd.Payment_MethodID,
		ShipmentID:       shopping_cartadd.ShipmentID,
		Quantity:         shopping_cartadd.Quantity,
	}
}
