package transaction_history_service

import (
	"toko-belanja-app/repository/product_repository"
	"toko-belanja-app/repository/transaction_history_repository"
	"toko-belanja-app/repository/user_repository"
)

type TransactionHistoryService interface {
}

type transactionHistoryServiceImpl struct {
	thr transaction_history_repository.TransactionHistoryRepository
	pr  product_repository.ProductRepository
	ur  user_repository.UserRepository
}

func NewTransactionHistoryService(transactionHistoryRepo transaction_history_repository.TransactionHistoryRepository, productRepo product_repository.ProductRepository, userRepo user_repository.UserRepository) TransactionHistoryService {
	return &transactionHistoryServiceImpl{
		thr: transactionHistoryRepo,
		pr:  productRepo,
		ur:  userRepo,
	}
}
