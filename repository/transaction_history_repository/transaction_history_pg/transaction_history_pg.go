package transaction_history_pg

import (
	"database/sql"
	"toko-belanja-app/repository/transaction_history_repository"
)

type transactionHistoryPg struct {
	db *sql.DB
}

func NewTransactionHistoryPg(db *sql.DB) transaction_history_repository.TransactionHistoryRepository {
	return &transactionHistoryPg{db: db}
}
