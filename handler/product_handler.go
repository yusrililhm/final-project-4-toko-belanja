package handler

import (
	"strconv"
	"toko-belanja-app/dto"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/service/product_service"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productHandlerImpl struct {
	ps product_service.ProductService
}

func NewProductHandler(productService product_service.ProductService) ProductHandler {
	return &productHandlerImpl{ps: productService}
}

// AddProduct implements ProductHandler.
// AddProduct godoc
// @Summary Add product
// @Description Add product
// @Tags Products
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param dto.ProductRequest body dto.ProductRequest true "body request for add product"
// @Success 201 {object} dto.ProductResponse
// @Router /products [post]
func (ph *productHandlerImpl) AddProduct(ctx *gin.Context) {
	
	addRequest := &dto.ProductRequest{}

	if err := ctx.ShouldBindJSON(addRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}
}

// DeleteProduct implements ProductHandler.
// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product
// @Tags Products
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param ProductId path int true "Product Id"
// @Success 200 {object} dto.ProductResponse
// @Router /products/{productId} [delete]
func (ph *productHandlerImpl) DeleteProduct(ctx *gin.Context) {
	productId, _ := strconv.Atoi(ctx.Param("productId"))
	_ = productId
}

// GetProducts implements ProductHandler.
// GetProducts godoc
// @Summary Get products
// @Description Get products
// @Tags Products
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Success 200 {object} dto.ProductResponse
// @Router /products [get]
func (ph *productHandlerImpl) GetProducts(ctx *gin.Context) {

}

// UpdateProduct implements ProductHandler.
// UpdateProduct godoc
// @Summary Update product
// @Description Update product
// @Tags Products
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param ProductId path int true "Product Id"
// @Param dto.ProductRequest body dto.ProductRequest true "body request for update product"
// @Success 200 {object} dto.ProductResponse
// @Router /products/{productId} [put]
func (ph *productHandlerImpl) UpdateProduct(ctx *gin.Context) {

	updateRequest := &dto.ProductRequest{}

	if err := ctx.ShouldBindJSON(updateRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	productId, _ := strconv.Atoi(ctx.Param("productId"))
	_ = productId
}
