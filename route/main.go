package route

import "githumb.com/Abeldlp/price-checker/config"

func InitializeRoutes() {

	// V1バージョンAPIルート設定
	InitializeProductRoutes(config.V1)
}
