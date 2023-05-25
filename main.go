package main

import (
	"log"

	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/cron"
	"githumb.com/Abeldlp/price-checker/service"
)

func main() {
	// Initialize Application
	config.InitilializeEnvironmentVariables()
	config.InitializeDatabase()
	scheduler := cron.InitializeCron()

	// Initialize Services
	proSer := service.NewProductService()
	usrSer := service.NewUserService()

	scheduler.Every(1).Second().Do(func() {
		products, err := proSer.GetAllProducts()
		if err != nil {
			log.Fatal(err)
		}

		for _, product := range *products {
			price, err := proSer.GetProductPrice(product.Url)
			if err != nil {
				log.Fatal(err)
			}

			if price < product.CurrentPrice {
				product.CurrentPrice = price
				user, err := proSer.GetProductUser(product.Id)
				if err != nil {
					log.Fatal(err)
				}

				usrSer.NotifyUser(user.Email)

			}
		}
	})

	go scheduler.StartBlocking()

	select {}
}
