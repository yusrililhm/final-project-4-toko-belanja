package product_service

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/product_repository"
)

type ProductService interface {
	CreateProduct(productPayLoad *dto.ProductRequest) (*dto.ProductResponse, errs.Error)
	GetAllProduct() (*dto.ProductResponse, errs.Error)
	UpdateProduct(productId int, productPayLoad *dto.ProductRequest) (*dto.ProductResponse, errs.Error)
	DeleteProduct(productId int) (*dto.ProductResponse, errs.Error)
}

type productServiceImpl struct {
	pr product_repository.ProductRepository
}

func NewProductService(ProductRepo product_repository.ProductRepository) ProductService {
	return &productServiceImpl{pr: ProductRepo}
}

func (ps *productServiceImpl) CreateProduct(productPayLoad *dto.ProductRequest) (*dto.ProductResponse, errs.Error)
func (ps *productServiceImpl) GetAllProduct() (*dto.ProductResponse, errs.Error)
func (ps *productServiceImpl) UpdateProduct(productId int, productPayLoad *dto.ProductRequest) (*dto.ProductResponse, errs.Error)
func (ps *productServiceImpl) DeleteProduct(productId int) (*dto.ProductResponse, errs.Error)
