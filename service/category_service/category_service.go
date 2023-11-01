package category_service

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/category_repository"
)

type CategoryService interface {
	CreateCategory(categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
	GetAllCategory() (*dto.CategoryResponse, errs.Error)
	UpdateCategory(categoryId int, categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
	DeleteCategory(categoryId int) (*dto.CategoryResponse, errs.Error)
}

type categoryServiceImpl struct {
	cr category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{cr: categoryRepo}
}

func (cs *categoryServiceImpl) CreateCategory(categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
func (cs *categoryServiceImpl) GetAllCategory() (*dto.CategoryResponse, errs.Error)
func (cs *categoryServiceImpl) UpdateCategory(categoryId int, categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
func (cs *categoryServiceImpl) DeleteCategory(categoryId int) (*dto.CategoryResponse, errs.Error)
