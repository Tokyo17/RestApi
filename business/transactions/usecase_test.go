package transactions_test

import (
	"context"
	"final/app/middleware"
	"final/business/transactions"
	"final/business/transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepository mocks.Repository
var transactionService transactions.UseCase
var transactionDomain transactions.TransactionDomain
var transactionDetailDomain transactions.Transaction_DetailDomain
var paymentMethodDomain transactions.Payment_MethodDomain
var shipmentDomain transactions.ShipmentDomain

var userToken middleware.ConfigJWT

func setup() {
	transactionService = transactions.NewTransactionUseCase(&transactionRepository, time.Hour*1, userToken)
	transactionDomain = transactions.TransactionDomain{
		ID:               1,
		Status:           "Dibayar",
		UserID:           1,
		ProductID:        1,
		Quantity:         1,
		Total_Price:      100000,
		Payment_MethodID: 1,

		ShipmentID: 1,
	}
	paymentMethodDomain = transactions.Payment_MethodDomain{
		ID:   1,
		Name: "OVO",
	}

	transactionDetailDomain = transactions.Transaction_DetailDomain{
		UserID:         1,
		StatusShipment: "Undelivered",
		TransactionID:  1,
		ProductID:      1,
	}

	shipmentDomain = transactions.ShipmentDomain{
		ID:             1,
		Name:           "JNE",
		Shipment_Price: 50000,
	}

}

func TestCheckout(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Add", func(t *testing.T) {
		transactionRepository.On("Add",
			mock.Anything,
			mock.Anything).Return(transactionDomain, nil).Once()
		_, err := transactionService.Add(context.Background(), transactions.TransactionDomain{
			ProductID:        1,
			Payment_MethodID: 1,
			ShipmentID:       1,
			Quantity:         1,
		})
		assert.Nil(t, err)
	})
}

func TestDetailSC(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid DetailSC", func(t *testing.T) {
		transactionRepository.On("DetailSC",
			mock.Anything,
			mock.AnythingOfType("int")).Return([]transactions.TransactionDomain{}, nil).Once()
		_, err := transactionService.DetailSC(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestGetTransDetail(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid GetTransDetail", func(t *testing.T) {
		transactionRepository.On("GetTransDetail",
			mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(transactionDetailDomain, transactionDomain, nil).Once()
		_, _, err := transactionService.GetTransDetail(context.Background(), 1, 1)
		assert.Nil(t, err)
	})
}

func TestAddPM(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid AddPM", func(t *testing.T) {
		transactionRepository.On("AddPM",
			mock.Anything,
			mock.Anything).Return(paymentMethodDomain, nil).Once()
		_, err := transactionService.AddPM(context.Background(), transactions.Payment_MethodDomain{
			Name: "JNE EXPRESS",
		})
		assert.Nil(t, err)
	})
}

func TestGetPM(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid GetPM", func(t *testing.T) {
		transactionRepository.On("GetPM",
			mock.Anything).Return([]transactions.Payment_MethodDomain{}, nil).Once()

		_, err := transactionService.GetPM(context.Background())
		assert.Nil(t, err)
	})
}

func TestGetShipment(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid GetPM", func(t *testing.T) {
		transactionRepository.On("GetShipment",
			mock.Anything).Return([]transactions.ShipmentDomain{}, nil).Once()

		_, err := transactionService.GetShipment(context.Background())
		assert.Nil(t, err)
	})
}

// AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error)
// GetTransDetail(ctx context.Context, userid, transid int) (Transaction_DetailDomain, TransactionDomain, error)
// DetailSC(ctx context.Context, id int) ([]TransactionDomain, error)
