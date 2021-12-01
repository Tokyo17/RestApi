package transactions

import (
	"final/app/middleware"
	"final/business/transactions"
	"final/controllers"

	"errors"
	requests "final/controllers/transactions/requets"
	"final/controllers/transactions/respons"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUseCase transactions.UseCase
}

func NewTransactionController(transactionUseCase transactions.UseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
	}
}

func (transactionController TransactionController) GetPM(c echo.Context) error {
	ctx := c.Request().Context()
	payment, err := transactionController.TransactionUseCase.GetPM(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNoContent, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListPMFromDomain(payment))
}
func (transactionController TransactionController) AddPM(c echo.Context) error {
	var payment_method requests.Payment_MethodAdd
	c.Bind(&payment_method)
	ctx := c.Request().Context()
	payment_methodAdd := payment_method.ToDomain()
	newPayment, err := transactionController.TransactionUseCase.AddPM(ctx, payment_methodAdd)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.PMFromDomain(newPayment))
}

func (transactionController TransactionController) AddShipment(c echo.Context) error {
	var shipment requests.ShipmentAdd
	c.Bind(&shipment)
	ctx := c.Request().Context()
	shipmentAdd := shipment.ToDomain()
	newShipment, err := transactionController.TransactionUseCase.AddShipment(ctx, shipmentAdd)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, newShipment)
}

func (transactionController TransactionController) GetShipment(c echo.Context) error {
	ctx := c.Request().Context()
	shipment, err := transactionController.TransactionUseCase.GetShipment(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNoContent, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListShipmentFromDomain(shipment))
}

func (transactioncontroller TransactionController) Add(c echo.Context) error {

	shopping_cart := requests.Shopping_CartAdd{}
	shopping_cart.UserID = middleware.GetClaimsUserId(c)
	c.Bind(&shopping_cart)

	ctx := c.Request().Context()
	transaction, err := transactioncontroller.TransactionUseCase.Add(ctx, shopping_cart.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.FromDomain(transaction))
}

func (transactionController TransactionController) DetailSC(c echo.Context) error {
	ctx := c.Request().Context()
	listSC, err := transactionController.TransactionUseCase.DetailSC(ctx, middleware.GetClaimsUserId(c))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNoContent, errors.New("invalid payment, please enter number same as total price"))
	}
	return controllers.NewSuccessResponse(c, respons.ListSCromDomain(listSC))
}

func (TransactionController TransactionController) Pay(c echo.Context) error {
	pay := requests.Payment{}
	c.Bind(&pay)
	transactionid, _ := strconv.Atoi(c.QueryParam("transactionid"))
	ctx := c.Request().Context()
	result, err := TransactionController.TransactionUseCase.Pay(ctx, transactionid, pay.Total_Price)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.FromDomain(result))

}

func (TransactionController TransactionController) GetTransDetail(c echo.Context) error {
	userid := middleware.GetClaimsUserId(c)
	transactionid, _ := strconv.Atoi(c.QueryParam("transactionid"))
	ctx := c.Request().Context()
	detail, transaction, err := TransactionController.TransactionUseCase.GetTransDetail(ctx, userid, transactionid)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNoContent, err)
	}
	return controllers.NewSuccessResponseDetails(c, respons.DetailFromDomain(detail), respons.ShoppingCartFromDomain(transaction))
}
