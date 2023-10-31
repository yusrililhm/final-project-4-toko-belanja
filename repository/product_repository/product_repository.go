package product_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type ProductRepository interface {
	CreateNewProduct(productPayLoad *entity.Product) (*dto.NewProductResponse, errs.Error)
}
