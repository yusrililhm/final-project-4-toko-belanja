package handler

import (
	"strconv"
	"toko-belanja-app/dto"
	"toko-belanja-app/pkg/errs"
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
// AddCategory godoc
// @Summary Add category
// @Description Add category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param dto.CategoriesRequest body dto.CategoriesRequest true "body request for add category"
// @Success 201 {object} dto.CategoryResponse
// @Router /categories [post]
func (ch *categoryHandlerImpl) AddCategory(ctx *gin.Context) {

	addRequest := &dto.CategoriesRequest{}

	if err := ctx.ShouldBindJSON(addRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}
}

// DeleteCategory implements CategoryHandler.
// DeleteCategory godoc
// @Summary Delete category
// @Description Delete category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param CategoryId path int true "Category Id"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories/{categoryId} [delete]
func (ch *categoryHandlerImpl) DeleteCategory(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))
	_ = categoryId
}

// GetCategories implements CategoryHandler.
// GetCategories godoc
// @Summary Get categories
// @Description Get categories
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories [get]
func (ch *categoryHandlerImpl) GetCategories(ctx *gin.Context) {

}

// UpdateCategory implements CategoryHandler.
// UpdateCategory godoc
// @Summary Update category
// @Description Update category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param CategoryId path int true "Category Id path"
// @Param dto.CategoriesRequest body dto.CategoriesRequest true "body request for update category"
// @Success 200 {object} dto.CategoryResponse
// @Router /categories/{categoryId} [patch]
func (ch *categoryHandlerImpl) UpdateCategory(ctx *gin.Context) {

	updateRequest := &dto.CategoriesRequest{}

	if err := ctx.ShouldBindJSON(updateRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))
	_ = categoryId
}
