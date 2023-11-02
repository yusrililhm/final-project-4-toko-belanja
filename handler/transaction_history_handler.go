package handler

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
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
// AddTransaction godoc
// @Summary Add transaction
// @Description Add transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param dto.TransactionRequest body dto.TransactionRequest true "body request for add transaction"
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} dto.TransactionHistoryResponse
// @Router /transactions [post]
func (th *transactionHistoryHandlerImpl) AddTransaction(ctx *gin.Context) {
	
	addRequest := &dto.TransactionRequest{}

	if err := ctx.ShouldBindJSON(addRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	userData := ctx.MustGet("userData").(entity.User)

	response, err := th.ths.CreateTransaction(userData.Id, addRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.Code, response)
}

// GetMyTransaction implements TransactionHistoryHandler.
// GetMyTransaction godoc
// @Summary Get my transaction
// @Description Get my transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.TransactionHistoryResponse
// @Router /transactions/my-transactions [get]
func (th *transactionHistoryHandlerImpl) GetMyTransaction(ctx *gin.Context) {
	response, err := th.ths.GetTransactionWithProducts()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.Code, response)
}

// GetUsersTransaction implements TransactionHistoryHandler.
// GetUsersTransaction godoc
// @Summary Get user transaction
// @Description Get user transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.TransactionHistoryResponse
// @Router /transactions/user-transactions [get]
func (th *transactionHistoryHandlerImpl) GetUsersTransaction(ctx *gin.Context) {
	response, err := th.ths.GetTransactionWithProductsAndUser()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.Code, response)
}
