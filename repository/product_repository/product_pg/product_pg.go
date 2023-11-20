package product_pg

import (
	"database/sql"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/product_repository"
)

const (
	createProduct = `
		INSERT INTO products (
			title, 
			price,
			stock,
			category_Id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id, title, price, stock, category_id, created_at
	`

	getProduct = `
		SELECT
			p.id,
			p.title,
			p.price,
			p.stock,
			p.category_id,
			p.created_at
		FROM
			products AS p
		WHERE
			p.deleted_at IS NULL
		ORDER BY
			p.id ASC
	`

	getProductById = `
		SELECT 
			p.id,
			p.title,
			p.price,
			p.stock,
			p.category_id,
			p.created_at
		FROM 
			products AS p
		WHERE 
			p.id = $1
			AND p.deleted_at IS NULL
	`

	updateProductById = `
		UPDATE
			products
		SET
			title = $2,
			price = $3,
			stock = $4,
			category_Id = $5,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, price, stock, category_id, created_at, updated_at
	`

	deleteProductById = `
		UPDATE
			products
		SET
			deleted_at = now()
		WHERE
			id = $1
	`
)

type productPg struct {
	db *sql.DB
}

func NewProductPg(db *sql.DB) product_repository.ProductRepository {
	return &productPg{db: db}
}

func (p *productPg) CreateNewProduct(productPayLoad *entity.Product) (*dto.NewProductResponse, errs.Error) {
	tx, err := p.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var product dto.NewProductResponse

	row := tx.QueryRow(
		createProduct,
		productPayLoad.Title,
		productPayLoad.Price,
		productPayLoad.Stock,
		productPayLoad.CategoryId,
	)
	err = row.Scan(
		&product.Id,
		&product.Title,
		&product.Price,
		&product.Stock,
		&product.CategoryId,
		&product.CreatedAt,
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

	return &product, nil
}

func (p *productPg) GetAllProducts() ([]entity.Product, errs.Error) {
	rows, err := p.db.Query(getProduct)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.Id,
			&product.Title,
			&product.Price,
			&product.Stock,
			&product.CategoryId,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return products, nil
}

func (p *productPg) GetProductById(id int) (*entity.Product, errs.Error) {
	var product entity.Product

	err := p.db.QueryRow(getProductById, id).Scan(
		&product.Id,
		&product.Title,
		&product.Price,
		&product.Stock,
		&product.CategoryId,
		&product.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}

func (p *productPg) UpdateProductById(productPayLoad *entity.Product) (*dto.UpdateProductResponse, errs.Error) {
	tx, err := p.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateProductById, productPayLoad.Id, productPayLoad.Title, productPayLoad.Price, productPayLoad.Stock, productPayLoad.CategoryId)

	var productUpdate dto.UpdateProductResponse
	err = row.Scan(
		&productUpdate.Id,
		&productUpdate.Title,
		&productUpdate.Price,
		&productUpdate.Stock,
		&productUpdate.CategoryId,
		&productUpdate.CreatedAt,
		&productUpdate.UpdatedAt,
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

	return &productUpdate, nil
}

func (p *productPg) DeleteProductById(productId int) errs.Error {
	tx, err := p.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteProductById, productId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
