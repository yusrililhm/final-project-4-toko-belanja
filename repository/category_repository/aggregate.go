package category_repository

import (
	"time"
	"toko-belanja-app/entity"
)

type CategoryProduct struct {
	Category entity.Category
	Product  entity.Product
}

type CategoryProductMapped struct {
	Id                int              `json:"id"`
	Type              string           `json:"type"`
	SoldProductAmount int              `json:"sold_product_amount"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
	Products          []entity.Product `json:"products"`
}

func (ctm *CategoryProductMapped) HandleMappingCategoryWithProduct(categoryProduct []*CategoryProduct) []*CategoryProductMapped {
	categoryProductsMapped := []*CategoryProductMapped{}

	for _, eachCategoryProduct := range categoryProduct {
		isCategoryExists := false

		for i := range categoryProductsMapped {
			if eachCategoryProduct.Category.Id == categoryProduct[i].Category.Id {
				isCategoryExists = true
				categoryProductsMapped[i].Products = append(categoryProductsMapped[i].Products, eachCategoryProduct.Product)
				break
			}
		}

		if !isCategoryExists {
			categoryProductMapped := &CategoryProductMapped{
				Id:                eachCategoryProduct.Category.Id,
				Type:              eachCategoryProduct.Category.Type,
				SoldProductAmount: eachCategoryProduct.Category.SoldProductAmount,
				CreatedAt:         eachCategoryProduct.Category.CreatedAt,
				UpdatedAt:         eachCategoryProduct.Category.UpdatedAt,
			}

			categoryProductMapped.Products = append(categoryProductMapped.Products, eachCategoryProduct.Product)

			if categoryProductMapped.Products[0].Id == 0 {
				categoryProductMapped.Products = []entity.Product{}
			}

			categoryProductsMapped = append(categoryProductsMapped, categoryProductMapped)
		}
	}

	return categoryProductsMapped
}
