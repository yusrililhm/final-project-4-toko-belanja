package handler

import (
	"toko-belanja-app/service/transaction_history_service"

	"github.com/gin-gonic/gin"
)

type TransactionHistoryHandler interface {
	AddTransaction(ctx *gin.Context)
	GetMyTransaction(ctx *gin.Context)
	GetUsersTransaction(ctx *gin.Context)
}

type transactionHistoryHandlerImpl struct {
	ths transaction_history_service.TransactionHistoryService
}

func NewTransactionHistoryHandler(transactionhistoryService transaction_history_service.TransactionHistoryService) TransactionHistoryHandler {
	return &transactionHistoryHandlerImpl{ths: transactionhistoryService}
}

// AddTransaction implements TransactionHistoryHandler.
func (th *transactionHistoryHandlerImpl) AddTransaction(ctx *gin.Context) {
	
}

// GetMyTransaction implements TransactionHistoryHandler.
func (th *transactionHistoryHandlerImpl) GetMyTransaction(ctx *gin.Context) {
	
}

// GetUsersTransaction implements TransactionHistoryHandler.
func (th *transactionHistoryHandlerImpl) GetUsersTransaction(ctx *gin.Context) {
	
}
