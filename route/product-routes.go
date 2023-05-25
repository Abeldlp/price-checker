package route

import (
	"github.com/gin-gonic/gin"
	"githumb.com/Abeldlp/price-checker/controller"
)

// InitializeProductRoutes 商品関連のルートを設定
func InitializeProductRoutes(group *gin.RouterGroup) {
	products := group.Group("/products")

	// 商品一覧取得
	products.POST("", controller.CreateProductNotification)
}
