package service

import (
	"fmt"
	"log"

	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/model"
)

type ProductService interface {
	UpdateProduct(product model.Product) bool
	GetProductPrice(productUrl string) (int, error)
	GetAllProducts() (*[]model.Product, error)
	GetProductUser(productId int) (*model.User, error)
}

type PService struct{}

func NewProductService() ProductService {
	return &PService{}
}

func (p *PService) UpdateProduct(product model.Product) bool {
	fmt.Println("Saving the following product:", product)
	qry := fmt.Sprintf("UPDATE products SET current_price = %d, url = '%s' WHERE id = %d", product.CurrentPrice, product.Url, product.Id)
	_, err := config.DB.Exec(qry)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (p *PService) GetProductPrice(productUrl string) (int, error) {
	return 2000, nil
}

func (p *PService) GetProductUser(productId int) (*model.User, error) {
	qry := fmt.Sprintf("SELECT u.* FROM products as p LEFT JOIN users as u ON p.user_id=u.id WHERE p.id = %d", productId)
	var user model.User

	rows, err := config.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return &user, nil
}

func (p *PService) GetAllProducts() (*[]model.Product, error) {
	qry := "SELECT * FROM products"
	var products []model.Product

	rows, err := config.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.CurrentPrice, &product.Url, &product.UserId)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return &products, nil
}
