package views

import (
	"github.com/BalusChen/IKHNAIE_API/handler/product"
	"github.com/BalusChen/IKHNAIE_API/handler/transaction"
	"github.com/BalusChen/IKHNAIE_API/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	r := e.Group("ikhnaie/v1/")

	initUserRoutes(r)
	initProductRoutes(r)
	initTransactionRoutes(r)
}

func initUserRoutes(r *gin.RouterGroup) {
	router := r.Group("user/")

	router.GET("info", user.GetInformation)
}

func initProductRoutes(r *gin.RouterGroup) {
	router := r.Group("product/")

	router.GET("info", product.GetInformation)
}

func initTransactionRoutes(r *gin.RouterGroup) {
	router := r.Group("transaction/")

	router.GET("info", transaction.GetInformation)
}
