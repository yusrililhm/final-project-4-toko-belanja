package category_service

import (
	"net/http"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/pkg/helpers"
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

func (cs *categoryServiceImpl) CreateCategory(categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error) {
	err := helpers.ValidateStruct(categoryPayLoad)

	if err != nil {
		return nil, err
	}

	category := &entity.Category{
		Type: categoryPayLoad.Type,
	}

	response, err := cs.cr.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		Code:    http.StatusCreated,
		Message: "Category has been successfully created",
		Data:    response,
	}, nil
}

func (cs *categoryServiceImpl) GetAllCategory() (*dto.CategoryResponse, errs.Error) {
	response, err := cs.cr.GetCategory()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.CategoryResponse{
		Code:    http.StatusOK,
		Message: "Category has been successfully fetched",
		Data:    response,
	}, nil
}

func (cs *categoryServiceImpl) UpdateCategory(categoryId int, categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error) {
	err := helpers.ValidateStruct(categoryPayLoad)

	if err != nil {
		return nil, err
	}

	checkCategory, err := cs.cr.CheckCategoryId(categoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	if checkCategory.Id != categoryId {
		return nil, errs.NewNotFoundError("invalid user")
	}

	category := &entity.Category{
		Id: categoryId,
		Type: categoryPayLoad.Type,
	}

	response, err := cs.cr.UpdateCategory(category)

	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		Code:    http.StatusOK,
		Message: "Category has been successfully updated",
		Data:    response,
	}, nil
}

func (cs *categoryServiceImpl) DeleteCategory(categoryId int) (*dto.CategoryResponse, errs.Error) {
	checkCategory, err := cs.cr.CheckCategoryId(categoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid user")
		}
		return nil, err
	}

	if checkCategory.Id != categoryId {
		return nil, errs.NewNotFoundError("invalid user")
	}

	cs.cr.DeleteCategory(categoryId)

	return &dto.CategoryResponse{
		Code: http.StatusOK,
		Message: "Category has been successfully deleted",
		Data: nil,
	}, nil
}
