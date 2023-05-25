package config

import (
	"github.com/gin-gonic/gin"
)

var Server *gin.Engine
var V1 *gin.RouterGroup

func InitializeGinServer() {
	// デフォルトのミドルウェアを使用してルーターを作成
	router := gin.Default()

	// ルートグループを作成
	V1 = router.Group("/api/v1")

	// サーバーをグローバル変数に格納
	Server = router
}
