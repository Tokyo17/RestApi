package requests

import "final/business/transactions"

type ShipmentAdd struct {
	Name           string  `json:"name"`
	Shipment_Price float64 `json:"shipment_price"`
}

func (shipmentAdd *ShipmentAdd) ToDomain() transactions.ShipmentDomain {
	return transactions.ShipmentDomain{
		Name:           shipmentAdd.Name,
		Shipment_Price: shipmentAdd.Shipment_Price,
	}
}
