package product_service

import (
	"net/http"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/pkg/helpers"
	"toko-belanja-app/repository/category_repository"
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
	cr category_repository.CategoryRepository
}

func NewProductService(productRepo product_repository.ProductRepository, categoryRepo category_repository.CategoryRepository) ProductService {
	return &productServiceImpl{
		pr: productRepo,
		cr: categoryRepo,
	}
}

func (ps *productServiceImpl) CreateProduct(productPayLoad *dto.ProductRequest) (*dto.ProductResponse, errs.Error) {
	err := helpers.ValidateStruct(productPayLoad)

	if err != nil {
		return nil, err
	}

	checkCategoryId, err := ps.cr.CheckCategoryId(productPayLoad.CategoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("Category not found")
		}
		return nil, err
	}

	if checkCategoryId.Id != productPayLoad.CategoryId {
		return nil, errs.NewNotFoundError("invalid category id")
	}

	product := &entity.Product{
		Title:      productPayLoad.Title,
		Price:      productPayLoad.Price,
		Stock:      productPayLoad.Stock,
		CategoryId: productPayLoad.CategoryId,
	}

	response, err := ps.pr.CreateNewProduct(product)

	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		Code:    http.StatusCreated,
		Message: "Your product has been successfully created",
		Data:    response,
	}, nil
}

func (ps *productServiceImpl) GetAllProduct() (*dto.ProductResponse, errs.Error) {
	response, err := ps.pr.GetAllProducts()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.ProductResponse{
		Code:    http.StatusOK,
		Message: "Products has been successfully fetched",
		Data:    response,
	}, nil
}

func (ps *productServiceImpl) UpdateProduct(productId int, productPayLoad *dto.ProductRequest) (*dto.ProductResponse, errs.Error) {
	err := helpers.ValidateStruct(productPayLoad)

	if err != nil {
		return nil, err
	}

	checkProductId, err := ps.pr.GetProductById(productId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("Product not found")
		}
		return nil, err
	}

	if checkProductId.Id != productId {
		return nil, errs.NewNotFoundError("invalid product")
	}

	checkCategoryId, err := ps.cr.CheckCategoryId(productPayLoad.CategoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("Category not found")
		}
		return nil, err
	}

	if checkCategoryId.Id != productPayLoad.CategoryId {
		return nil, errs.NewNotFoundError("invalid category id")
	}

	product := &entity.Product{
		Id:         productId,
		Title:      productPayLoad.Title,
		Price:      productPayLoad.Price,
		Stock:      productPayLoad.Stock,
		CategoryId: productPayLoad.CategoryId,
	}

	response, err := ps.pr.UpdateProductById(product)

	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		Code:    http.StatusOK,
		Message: "Product has been successfully updated",
		Data:    response,
	}, nil
}

func (ps *productServiceImpl) DeleteProduct(productId int) (*dto.ProductResponse, errs.Error) {
	checkProductId, err := ps.pr.GetProductById(productId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("Product not found")
		}
		return nil, err
	}

	if checkProductId.Id != productId {
		return nil, errs.NewNotFoundError("invalid product")
	}

	err = ps.pr.DeleteProductById(productId)

	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		Code:    http.StatusOK,
		Message: "Product has been successfully deleted",
		Data:    nil,
	}, nil
}
