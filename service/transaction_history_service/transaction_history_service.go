package transaction_history_service

import (
	"net/http"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/pkg/helpers"
	"toko-belanja-app/repository/category_repository"
	"toko-belanja-app/repository/product_repository"
	"toko-belanja-app/repository/transaction_history_repository"
	"toko-belanja-app/repository/user_repository"
)

type TransactionHistoryService interface {
	CreateTransaction(productId int, transactionPayLoad *dto.TransactionRequest) (*dto.TransactionHistoryResponse, errs.Error)
	GetTransactionWithProducts(userId int) (*dto.TransactionHistoryResponse, errs.Error)
	GetTransactionWithProductsAndUser() (*dto.TransactionHistoryResponse, errs.Error)
}

type transactionHistoryServiceImpl struct {
	thr transaction_history_repository.TransactionHistoryRepository
	pr  product_repository.ProductRepository
	ur  user_repository.UserRepository
	cr	category_repository.CategoryRepository
}

func NewTransactionHistoryService(transactionHistoryRepo transaction_history_repository.TransactionHistoryRepository, productRepo product_repository.ProductRepository, userRepo user_repository.UserRepository) TransactionHistoryService {
	return &transactionHistoryServiceImpl{
		thr: transactionHistoryRepo,
		pr:  productRepo,
		ur:  userRepo,
	}
}

func (ts *transactionHistoryServiceImpl) CreateTransaction(userId int, transactionPayLoad *dto.TransactionRequest) (*dto.TransactionHistoryResponse, errs.Error) {
	err := helpers.ValidateStruct(transactionPayLoad)

	if err != nil {
		return nil, err
	}

	product, err := ts.pr.GetProductById(transactionPayLoad.ProductId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	if transactionPayLoad.Quantity > product.Stock {
		return nil, errs.NewUnprocessableEntityError("Insufficient stock for the requested quantity")
	}

	user, err := ts.ur.GetUserById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	totalCost := product.Price * transactionPayLoad.Quantity
	if user.Balance < totalCost {
		return nil, errs.NewUnprocessableEntityError("Insufficient balance for the transaction")
	}

	transaction := &entity.TransactionHistory{
		UserId: userId,
		ProductId: transactionPayLoad.ProductId,
		Quantity: transactionPayLoad.Quantity,
	}

	response, err := ts.thr.CreateNewTransaction(transaction)

	if err != nil {
		return nil, err
	}

	return &dto.TransactionHistoryResponse{
		Code: http.StatusCreated,
		Message: "You have successfully purchased the product",
		Data: dto.TransactionBill{
			TotalPrice: response.TotalPrice,
			Quantity: response.Quantity,
			ProductTitle: response.ProductTitle,
		},
	}, nil
}
func (ts *transactionHistoryServiceImpl) GetTransactionWithProducts(userId int) (*dto.TransactionHistoryResponse, errs.Error)
func (ts *transactionHistoryServiceImpl) GetTransactionWithProductsAndUser() (*dto.TransactionHistoryResponse, errs.Error)
