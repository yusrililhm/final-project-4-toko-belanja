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
	VALUES ($1, $2, $3, ((SELECT p.price
	FROM products AS p
	WHERE
	id =$2)*$3))
	RETURNING
	product_id, quantity, p.title
	`
	getTransaction = `
		SELECT
			t.id,
			t.product_id,
			t.user_id,
			t.quantity,
			t.total_price,
			p.id AS product_id,
			p.title AS product_title,
			p.price,
			p.stock,
			p.category_id,
			p.created_at AS product_created_at,
			p.updated_at AS product_updated_at,
			u.id AS user_id,
			u.email,
			u.full_name,
			u.balance,
			u.created_at AS user_created_at,
			u.updated_at AS user_updated_at
		FROM
			transaction_histories AS t
		LEFT JOIN
			products AS p
		ON
			t.product_id = p.id
		LEFT JOIN
			users AS u
		ON
			t.user_id = u.id
		WHERE
			t.deleted_at IS NULL
		ORDER BY
			t.id ASC;
	`

	getMyTransaction = `
		SELECT t.id, 
			t.product_id,
			t.user_id, 
			t.quantity, 
			t.total_price, 
			p.id, 
			p.title, 
			p.price, 
			p.stock, 
			p.category_Id, 
			p.created_at, 
			p.updated_at
		FROM 
			transaction_histories as t
		LEFT JOIN
			products as p
		ON
			t.product_id = p.id
		WHERE t.user_id = $1 AND
		t.deleted_at IS NULL
		ORDER BY 
		t.id ASC
		`
)

func (t *transactionHistoryPg) CreateNewTransaction(transactionPayLoad *entity.TransactionHistory) (*dto.TransactionBill, errs.Error) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
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
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &transaction, nil
}

func (t *transactionHistoryPg) GetTransaction() ([]transaction_history_repository.TransactionProductMapped, errs.Error) {
	transactionProducts := []transaction_history_repository.TransactionProduct{}
	rows, err := t.db.Query(getTransaction)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		var transactionProduct transaction_history_repository.TransactionProduct

		err := rows.Scan(
			&transactionProduct.TransactionHistory.Id,
			&transactionProduct.TransactionHistory.ProductId,
			&transactionProduct.TransactionHistory.UserId,
			&transactionProduct.TransactionHistory.Quantity,
			&transactionProduct.TransactionHistory.TotalPrice,
			&transactionProduct.Product.Id,
			&transactionProduct.Product.Title,
			&transactionProduct.Product.Price,
			&transactionProduct.Product.Stock,
			&transactionProduct.Product.CategoryId,
			&transactionProduct.Product.CreatedAt,
			&transactionProduct.Product.UpdatedAt,
			&transactionProduct.User.Id,
			&transactionProduct.User.Email,
			&transactionProduct.User.FullName,
			&transactionProduct.User.Balance,
			&transactionProduct.User.CreatedAt,
			&transactionProduct.User.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		transactionProducts = append(transactionProducts, transactionProduct)
	}

	result := transaction_history_repository.TransactionProductMapped{}
	return result.HandleMappingTransactionWithProduct(transactionProducts), nil
}

func (t *transactionHistoryPg) GetMyTransaction(UserId int) ([]transaction_history_repository.MyTransactionProductMapped, errs.Error) {
	mytransactionProducts := []transaction_history_repository.MyTransactionProduct{}
	rows, err := t.db.Query(getMyTransaction, UserId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		var mytransactionProduct transaction_history_repository.MyTransactionProduct

		err := rows.Scan(
			&mytransactionProduct.TransactionHistory.Id,
			&mytransactionProduct.TransactionHistory.ProductId,
			&mytransactionProduct.TransactionHistory.UserId,
			&mytransactionProduct.TransactionHistory.Quantity,
			&mytransactionProduct.TransactionHistory.TotalPrice,
			&mytransactionProduct.Product.Id,
			&mytransactionProduct.Product.Title,
			&mytransactionProduct.Product.Price,
			&mytransactionProduct.Product.Stock,
			&mytransactionProduct.Product.CategoryId,
			&mytransactionProduct.Product.CreatedAt,
			&mytransactionProduct.Product.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		mytransactionProducts = append(mytransactionProducts, mytransactionProduct)
	}

	result := transaction_history_repository.MyTransactionProductMapped{}
	return result.HandleMappingMyTransactionWithProduct(mytransactionProducts), nil
}
