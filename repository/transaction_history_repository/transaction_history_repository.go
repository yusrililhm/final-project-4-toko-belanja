package transaction_history_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type TransactionHistoryRepository interface {
	CreateNewTransaction(transactionPayLoad *entity.TransactionHistory) (*dto.TransactionBill, errs.Error)
	GetMyTransaction(UserId int) ([]MyTransactionProductMapped, errs.Error)
	GetTransaction() ([]TransactionProductMapped, errs.Error)
}
