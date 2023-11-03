package category_pg

import (
	"database/sql"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/category_repository"
)

const (
	createNewCategory = `
		INSERT INTO "categories" 
		(
			type,
			sold_product_amount,
		)
		VALUES ($1, 0)
		RETURNING
			id, type, sold_product_amount, created_at;
	`
	updateCategoryById = `
		UPDATE
			categories
		SET
			type = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, type, sold_product_amount, updated_at
	`
	getCategoryWithProduct = `
			SELECT
			c.id,
			c.type,
			c.sold_product_amount,
			c.created_at,
			c.updated_at,
			p.id,
			p.title,
			p.price,
			p.stock,
			p.created_at,
			p.updated_at
		FROM
			categories as c
		LEFT JOIN
			products as p 
		ON 
			c.id = p.category_id
		WHERE
			c.deleted_at IS NULL
		ORDER BY
			c.id ASC
	`

	deleteCategoryById = `
		UPDATE
			categories
		SET
			deleted_at = now()
		WHERE
			id = $1
	`

	checkCategoryId = `
		SELECT 
			c.id,
			c.type,
			c.sold_product_amount,
		FROM 
			categories AS c
		WHERE
			c.id = $1
			AND c.deleted_at IS NULL
	`
)

type categoryPg struct {
	db *sql.DB
}

func NewCategoryPg(db *sql.DB) category_repository.CategoryRepository {
	return &categoryPg{db: db}
}

func (c *categoryPg) CreateCategory(categoryPayLoad *entity.Category) (*dto.CreateNewCategoriesResponse, errs.Error) {
	tx, err := c.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var category dto.CreateNewCategoriesResponse

	row := tx.QueryRow(createNewCategory, categoryPayLoad.Type)
	err = row.Scan(&category.Id, &category.Type, &category.SoldProductAmount, &category.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPg) GetCategory() ([]category_repository.CategoryProductMapped, errs.Error) {
	categoryProducts := []category_repository.CategoryProduct{}
	rows, err := c.db.Query(getCategoryWithProduct)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	for rows.Next() {
		categoryProduct := category_repository.CategoryProduct{}

		err := rows.Scan(
			&categoryProduct.Category.Id,
			&categoryProduct.Category.Type,
			&categoryProduct.Category.SoldProductAmount,
			&categoryProduct.Category.UpdatedAt,
			&categoryProduct.Category.CreatedAt,
			&categoryProduct.Product.Id,
			&categoryProduct.Product.Title,
			&categoryProduct.Product.Price,
			&categoryProduct.Product.Stock,
			&categoryProduct.Product.CreatedAt,
			&categoryProduct.Product.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong" + err.Error())
		}

		categoryProducts = append(categoryProducts, categoryProduct)
	}

	result := category_repository.CategoryProductMapped{}
	return result.HandleMappingCategoryWithProduct(categoryProducts), nil
}

func (c *categoryPg) UpdateCategory(categoryPayLoad *entity.Category) (*dto.UpdateCategoryResponse, errs.Error) {
	tx, err := c.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	row := tx.QueryRow(updateCategoryById, categoryPayLoad.Id, categoryPayLoad.Type)

	var categoryUpdate dto.UpdateCategoryResponse
	err = row.Scan(
		&categoryUpdate.Id,
		&categoryUpdate.Type,
		&categoryPayLoad.SoldProductAmount,
		&categoryUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &categoryUpdate, nil
}

func (c *categoryPg) CheckCategoryId(categoryId int) (*entity.Category, errs.Error) {
	category := entity.Category{}
	row := c.db.QueryRow(checkCategoryId, categoryId)
	err := row.Scan(&category.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewInternalServerError("rows not found" + err.Error())
		}
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &category, nil
}

func (c *categoryPg) DeleteCategory(categoryId int) errs.Error {
	tx, _ := c.db.Begin()

	_, err := tx.Exec(deleteCategoryById, categoryId)

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
