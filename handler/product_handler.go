package handler

import (
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
func (ph *productHandlerImpl) AddProduct(ctx *gin.Context) {

}

// DeleteProduct implements ProductHandler.
func (ph *productHandlerImpl) DeleteProduct(ctx *gin.Context) {

}

// GetProducts implements ProductHandler.
func (ph *productHandlerImpl) GetProducts(ctx *gin.Context) {

}

// UpdateProduct implements ProductHandler.
func (ph *productHandlerImpl) UpdateProduct(ctx *gin.Context) {

}
