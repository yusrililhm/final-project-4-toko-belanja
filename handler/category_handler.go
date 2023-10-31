package handler

import (
	"toko-belanja-app/service/category_service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	AddCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

type categoryHandlerImpl struct {
	cs category_service.CategoryService
}

func NewCategoryHandler(categoryService category_service.CategoryService) CategoryHandler {
	return &categoryHandlerImpl{cs: categoryService}
}

// AddCategory implements CategoryHandler.
func (ch *categoryHandlerImpl) AddCategory(ctx *gin.Context) {
	
}

// DeleteCategory implements CategoryHandler.
func (ch *categoryHandlerImpl) DeleteCategory(ctx *gin.Context) {
	
}

// GetCategories implements CategoryHandler.
func (ch *categoryHandlerImpl) GetCategories(ctx *gin.Context) {
	
}

// UpdateCategory implements CategoryHandler.
func (ch *categoryHandlerImpl) UpdateCategory(ctx *gin.Context) {
	
}
