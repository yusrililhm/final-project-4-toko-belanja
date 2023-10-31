package product_service

import "toko-belanja-app/repository/product_repository"

type ProductService interface {
}

type productServiceImpl struct {
	pr product_repository.ProductRepository
}

func NewProductService(ProductRepo product_repository.ProductRepository) ProductService {
	return &productServiceImpl{pr: ProductRepo}
}
