package transactions

import (
	"context"
	"errors"
	"final/app/middleware"
	"time"
)

type TransactionUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JwtToken       middleware.ConfigJWT
}

func NewTransactionUseCase(repo Repository, timeOut time.Duration, token middleware.ConfigJWT) UseCase {
	return &TransactionUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
		JwtToken:       token,
	}
}

func (uc *TransactionUseCase) Add(ctx context.Context, domain TransactionDomain) (TransactionDomain, error) {
	if domain.ProductID == 0 {
		return TransactionDomain{}, errors.New("product id is empty")
	}
	if domain.Quantity <= 0 {
		return TransactionDomain{}, errors.New("invalid quantity")
	}
	transaction, err := uc.Repo.Add(ctx, domain)
	if err != nil {
		return TransactionDomain{}, err
	}
	return transaction, nil
}

func (uc *TransactionUseCase) DetailSC(ctx context.Context, id int) ([]TransactionDomain, error) {
	transactions, err := uc.Repo.DetailSC(ctx, id)
	if err != nil {
		return nil, errors.New("invalid payment, please enter number same as total price")
	}
	return transactions, nil
}

func (uc *TransactionUseCase) GetPM(ctx context.Context) ([]Payment_MethodDomain, error) {
	payment_method, err := uc.Repo.GetPM(ctx)
	if err != nil {
		return nil, err
	}
	return payment_method, nil

}

func (uc *TransactionUseCase) AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error) {
	if domain.Name == "" {
		return Payment_MethodDomain{}, errors.New("payment method name is empty")
	}
	paymentMethod, err := uc.Repo.AddPM(ctx, domain)
	if err != nil {
		return Payment_MethodDomain{}, err
	}
	return paymentMethod, nil
}

func (uc *TransactionUseCase) AddShipment(ctx context.Context, domain ShipmentDomain) (ShipmentDomain, error) {
	if domain.Name == "" {
		domain.Name = "J&T"
	}

	if domain.Shipment_Price == 0 {
		domain.Shipment_Price = 50000
	}
	shipment, err := uc.Repo.AddShipment(ctx, domain)
	if err != nil {
		return ShipmentDomain{}, err
	}
	return shipment, nil
}

func (uc *TransactionUseCase) GetShipment(ctx context.Context) ([]ShipmentDomain, error) {
	shipment, err := uc.Repo.GetShipment(ctx)
	if err != nil {
		return nil, err
	}
	return shipment, nil
}

func (uc *TransactionUseCase) Pay(ctx context.Context, transactionid int, amount float64) (TransactionDomain, error) {
	pay, err := uc.Repo.Pay(ctx, transactionid, amount)
	if err != nil {
		return TransactionDomain{}, err
	}
	return pay, nil
}

func (uc *TransactionUseCase) GetTransDetail(ctx context.Context, userid, transid int) (Transaction_DetailDomain, TransactionDomain, error) {
	detail, trans, err := uc.Repo.GetTransDetail(ctx, userid, transid)
	if err != nil {
		return Transaction_DetailDomain{}, TransactionDomain{}, err
	}
	return detail, trans, nil
}
