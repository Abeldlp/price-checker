package main

import (
	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/cron"
	"githumb.com/Abeldlp/price-checker/route"
)

func main() {
	// 環境変数読み込み
	config.InitilializeEnvironmentVariables()

	// DB起動
	config.InitializeDatabase()
	defer config.CloseDatabase()

	// HTTP サーバー起動
	config.InitializeGinServer()
	route.InitializeRoutes()

	// スケジュール起動
	cron.InitializeCron()
	cron.NotifyUsersCronJob()

	// Goroutine起動
	go cron.Scheduler.StartBlocking()
	go config.Server.Run()

	// 終了待ち
	select {}
}
