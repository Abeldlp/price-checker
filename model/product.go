package model

import "fmt"

// Product 商品
type Product struct {
	Id           int    `json:"id"`
	CurrentPrice int    `json:"current_price"`
	Url          string `json:"url"`
	UserId       int    `json:"user_id"`
}

// NewProduct 商品を作成
func NewProduct(url string) *Product {
	return &Product{
		Url: url,
	}
}

// Scrape 商品の値段を取得
func (p *Product) Scrape() {
	p.CurrentPrice = 23
	fmt.Println(p.CurrentPrice)
}
