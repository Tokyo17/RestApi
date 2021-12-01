package products

import (
	"errors"
	"final/business/products"
	"final/controllers"
	"final/controllers/products/requests"
	"final/controllers/products/respons"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//STRUCT YANG BERISI INTERFACE USECASE DARI DOMAIN
type ProductController struct {
	ProductUseCase products.UseCase
}

// MENGISI STRUCT productController
func NewProductController(productUseCase products.UseCase) *ProductController {
	return &ProductController{
		ProductUseCase: productUseCase,
	}
}

//Memanggil Method Get untuk mendapatkan data dari domain
// berupa []ProductDomain dan error dan akan meretrun data
// berupa json dari baseRespont
func (ProductController ProductController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.Get(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNoContent, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListFromDomain(product))
}

func (ProductController ProductController) UploadType(c echo.Context) error {
	newProductType := requests.ProductTypeUpload{}
	c.Bind(&newProductType)
	uploadType := newProductType.ToDomain()
	ctx := c.Request().Context()
	productType, err := ProductController.ProductUseCase.UploadType(ctx, uploadType)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.TypeFromDomain(productType))

}

func (ProductController ProductController) FilterByType(c echo.Context) error {

	filterProduct, fail := strconv.Atoi(c.Param("id"))
	if fail != nil {
		return errors.New("gagal konversi product type id")
	}

	ctx := c.Request().Context()
	result, err := ProductController.ProductUseCase.FilterByType(ctx, filterProduct)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, respons.ListFromDomain(result))
}

func (ProductController ProductController) UploadProduct(c echo.Context) error {
	newProduct := requests.ProductUpload{}
	c.Bind(&newProduct)
	uploadProduct := newProduct.ToDomain()
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.UploadProduct(ctx, uploadProduct)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.ProductFromDomain(product))
}

func (ProductController ProductController) UpdateProduct(c echo.Context) error {
	newproduct := requests.ProductUpdate{}
	c.Bind(&newproduct)
	updateProduct := newproduct.ToDomain()
	ctx := c.Request().Context()
	product, err := ProductController.ProductUseCase.UpdateProduct(ctx, updateProduct, newproduct.ID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.ProductFromDomain(product))
}
