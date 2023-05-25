package cron

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"githumb.com/Abeldlp/price-checker/model"
	"githumb.com/Abeldlp/price-checker/service"
)

var Scheduler *gocron.Scheduler
var NotifyUserIntervalSeconds int = 1

// クローンサービスを起動
func InitializeCron() {
	s := gocron.NewScheduler(time.UTC)

	Scheduler = s
}

// 全てのプロダクトを集めて、プロダクトごとに値段の差を比較。前より低ければユーザに知らせる仕組み。
func NotifyUsersCronJob() {

	// サービス起動
	proSer := service.NewProductService()
	usrSer := service.NewUserService()

	Scheduler.Every(NotifyUserIntervalSeconds).Second().Do(func() {

		// 全てのプロダクトを集める
		products, err := proSer.GetAllProducts()
		if err != nil {
			log.Fatal(err)
		}

		// プロダクトごとにGoroutineで値段比較
		for _, product := range *products {
			go func(product model.Product) {
				price, err := proSer.GetProductPrice(product.Url)
				if err != nil {
					log.Fatal(err)
				}

				// 前より値段が低い場合新しい値段を保存し、登録ユーザにお知らせ
				if price < product.CurrentPrice {
					product.CurrentPrice = price
					proSer.UpdateProduct(product)
					user, err := proSer.GetProductUser(product.Id)
					if err != nil {
						log.Fatal(err)
					}

					usrSer.NotifyUser(user.Email)
				}
			}(product)
		}
	})
}
