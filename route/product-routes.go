package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitializeProductRoutes 商品関連のルートを設定
func InitializeProductRoutes(group *gin.RouterGroup) {
	products := group.Group("/products")

	// 商品一覧取得
	products.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
}
