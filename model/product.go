package model

import "fmt"

type Product struct {
	Id           int    `json:"id"`
	CurrentPrice int    `json:"current_price"`
	Url          string `json:"url"`
	UserId       int    `json:"user_id"`
}

func NewProduct(url string) *Product {
	return &Product{
		Url: url,
	}
}

func (p *Product) Scrape() {
	p.CurrentPrice = 23
	fmt.Println(p.CurrentPrice)
}
