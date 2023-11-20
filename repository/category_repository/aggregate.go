package category_repository

import (
	"time"
	"toko-belanja-app/entity"
)

type product struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Price     uint      `json:"price"`
	Stock     uint      `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryProduct struct {
	Category entity.Category
	Product  entity.Product
}

type CategoryProductMapped struct {
	Id                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Products          []product `json:"Products"`
}

func (ctm *CategoryProductMapped) HandleMappingCategoryWithProduct(categoryProduct []CategoryProduct) []CategoryProductMapped {
	categoryProductsMapped := make(map[int]CategoryProductMapped)

	for _, eachCategoryProduct := range categoryProduct {
		categoryId := eachCategoryProduct.Category.Id
		categoryProductMapped, exists := categoryProductsMapped[categoryId]
		if !exists {
			categoryProductMapped = CategoryProductMapped{
				Id:                eachCategoryProduct.Category.Id,
				Type:              eachCategoryProduct.Category.Type,
				SoldProductAmount: eachCategoryProduct.Category.SoldProductAmount,
				CreatedAt:         eachCategoryProduct.Category.CreatedAt,
				UpdatedAt:         eachCategoryProduct.Category.UpdatedAt,
			}
		}

		product := product{
			Id:        eachCategoryProduct.Product.Id,
			Title:     eachCategoryProduct.Product.Title,
			Price:     eachCategoryProduct.Product.Price,
			Stock:     eachCategoryProduct.Product.Stock,
			CreatedAt: eachCategoryProduct.Product.CreatedAt,
			UpdatedAt: eachCategoryProduct.Product.UpdatedAt,
		}
		categoryProductMapped.Products = append(categoryProductMapped.Products, product)
		categoryProductsMapped[categoryId] = categoryProductMapped
	}

	categoryProducts := []CategoryProductMapped{}
	for _, categoryProduct := range categoryProductsMapped {
		categoryProducts = append(categoryProducts, categoryProduct)
	}

	return categoryProducts
}
