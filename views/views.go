package views

import (
	"net/http"

	"github.com/BalusChen/IKHNAIE_API/handler/product"
	"github.com/BalusChen/IKHNAIE_API/handler/qrcode"
	"github.com/BalusChen/IKHNAIE_API/handler/transaction"
	"github.com/BalusChen/IKHNAIE_API/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	r := e.Group("ikhnaie/v1/")

	r.GET("ping", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.JSON(http.StatusOK, gin.H{
			"cartoon": "Tom and Jerry",
			"names":   []string{"Tom", "Jerry", "Sam"},
		})
	})

	initUserRoutes(r)
	initAdminRoutes(r)
	initProductRoutes(r)
	initTransactionRoutes(r)
	initQRCodeRoutes(r)
}

func initUserRoutes(r *gin.RouterGroup) {
	router := r.Group("user/")

	router.GET("info", user.Info)
	router.GET("list", user.List)
	router.POST("register", user.Register)
	router.POST("login", user.Login)
	router.GET("check", user.Check)
}

func initAdminRoutes(r *gin.RouterGroup) {
	router := r.Group("admin/")

	router.GET("info")
}

func initProductRoutes(r *gin.RouterGroup) {
	router := r.Group("product/")

	router.GET("info", product.GetInformation)
	router.POST("add", product.AddProduct)
}

func initTransactionRoutes(r *gin.RouterGroup) {
	router := r.Group("transaction/")

	router.GET("history", transaction.GetHistory)
	router.GET("add", transaction.AddTransaction)
}

func initQRCodeRoutes(r *gin.RouterGroup) {
	router := r.Group("qrcode/")

	router.GET("generate", qrcode.Generate)
	router.GET("retrieve", qrcode.Retrieve)
}
