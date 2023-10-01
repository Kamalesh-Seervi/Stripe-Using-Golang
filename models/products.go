package models

import "errors"

func GetAllProducts() ([]Product, error) {
	var products []Product
	tx := db.Find(&products)

	if tx.Error != nil {
		panic(tx.Error)
	}
	return products, nil
}

func CreateProduct(product []Product) error {
	tx := db.Create(product)
	return tx.Error
}

func GetProductById(id uint64) (Product, error) {
	var product Product
	tx := db.Where("id = ?", id).First(&product)
	if tx.Error != nil {
		return Product{}, tx.Error
	}

	if product.Id == 0 {
		return Product{}, errors.New("URL not found")
	}

	return product, nil
}

