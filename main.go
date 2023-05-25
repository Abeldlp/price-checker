package main

import (
	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/cron"
)

func main() {
	// アプリ起動
	config.InitilializeEnvironmentVariables()
	config.InitializeDatabase()
	cron.InitializeCron()

	// スケジュール起動
	cron.NotifyUsersCronJob()

	// クローンのブロック
	go cron.Scheduler.StartBlocking()

	select {}
}
