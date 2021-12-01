package requests

import "final/business/transactions"

type Payment_MethodAdd struct {
	Name string `json:"name"`
}

func (payment_methodAdd *Payment_MethodAdd) ToDomain() transactions.Payment_MethodDomain {
	return transactions.Payment_MethodDomain{
		Name: payment_methodAdd.Name,
	}
}
