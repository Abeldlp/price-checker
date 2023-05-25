package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githumb.com/Abeldlp/price-checker/model"
	"githumb.com/Abeldlp/price-checker/service"
)

type ProductRequest struct {
	ProductUrl string `json:"product_url"`
	UserEmail  string `json:"user_email"`
}

func CreateProductNotification(c *gin.Context) {
	var request ProductRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	productService := service.NewProductService()
	userService := service.NewUserService()

	// ユーザを保存
	userId := userService.SaveUser(request.UserEmail)

	// 商品を作成
	product := model.NewProduct(request.ProductUrl)
	product.UserId = userId

	// 商品を保存
	productSaved := productService.SaveProduct(product)
	if !productSaved {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}
