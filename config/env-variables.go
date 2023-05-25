package config

import (
	"log"

	"github.com/joho/godotenv"
)

// InitilializeEnvironmentVariables 環境変数を読み込む
func InitilializeEnvironmentVariables() {

	// .envファイルを読み込む
	err := godotenv.Load()

	// エラー処理
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
