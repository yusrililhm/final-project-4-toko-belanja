package product_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type ProductRepository interface {
	CreateNewProduct(productPayLoad *entity.Product) (*dto.NewProductResponse, errs.Error)
	GetAllProducts() (*[]entity.Product, errs.Error)
	GetProductById(id int) (*entity.Product, errs.Error)
	UpdateProductById(productPayLoad *entity.Product) (*dto.UpdateProductResponse, errs.Error)
	DeleteProductById(productId int) errs.Error
}
