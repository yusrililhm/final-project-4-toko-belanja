package transaction_history_pg

import (
	"database/sql"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/transaction_history_repository"
)

type transactionHistoryPg struct {
	db *sql.DB
}

func NewTransactionHistoryPg(db *sql.DB) transaction_history_repository.TransactionHistoryRepository {
	return &transactionHistoryPg{db: db}
}

const (
	createTransaction = `
		INSERT INTO transaction_histories (user_id, product_id, quantity, total_price)
		SELECT
			$1,
			$2,
			$3,
			(p.price * $3) AS total_price
		FROM products p
		WHERE p.id = $1
		RETURNING
			total_price,
			quantity,
			p.title AS product_title;
	`
)

func (t *transactionHistoryPg) CreateNewTransaction(transactionPayLoad *entity.TransactionHistory) (*dto.TransactionBill, errs.Error) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	var transaction dto.TransactionBill

	row := tx.QueryRow(
		createTransaction,
		transactionPayLoad.UserId,
		transactionPayLoad.ProductId,
		transactionPayLoad.Quantity,
	)
	err = row.Scan(
		&transaction.TotalPrice,
		&transaction.Quantity,
		&transaction.ProductTitle,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &transaction, nil
}
