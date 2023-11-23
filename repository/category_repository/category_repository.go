package category_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(categoryPayLoad *entity.Category) (*dto.CreateNewCategoriesResponse, errs.Error)
	GetCategory() ([]*CategoryProductMapped, errs.Error)
	UpdateCategory(categoryPayLoad *entity.Category) (*dto.UpdateCategoryResponse, errs.Error)
	CheckCategoryId(categoryId int) (*entity.Category, errs.Error)
	DeleteCategory(categoryId int) errs.Error
}
