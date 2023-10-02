package models

import "errors"

func GetAllProducts() ([]Product, error) {
	var products []Product
	tx := Dbase.Find(&products)

	if tx.Error != nil {
		panic(tx.Error)
	}
	return products, nil
}

func CreateProduct(product Product) error {
	tx := Dbase.Create(&product)
	return tx.Error
}

func GetProductById(product *Product, id string) (Product, error) {
	tx := Dbase.Where("id = ?", id).First(product)
	if tx.Error != nil {
		return *product, tx.Error
	}

	if product.Id == 0 {
		return *product, errors.New("Product not found")
	}

	return *product, nil
}
