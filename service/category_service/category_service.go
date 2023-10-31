package category_service

import "toko-belanja-app/repository/category_repository"

type CategoryService interface {
	
}

type categoryServiceImpl struct {
	cr category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{cr: categoryRepo}
}
