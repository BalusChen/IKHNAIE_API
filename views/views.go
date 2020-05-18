package views

import (
	"github.com/BalusChen/IKHNAIE_API/handler/blockchain"
	"net/http"

	"github.com/BalusChen/IKHNAIE_API/handler/product"
	"github.com/BalusChen/IKHNAIE_API/handler/qrcode"
	"github.com/BalusChen/IKHNAIE_API/handler/transaction"
	"github.com/BalusChen/IKHNAIE_API/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	r := e.Group("ikhnaie/v1/")

	initUserRoutes(r)
	initAdminRoutes(r)
	initProductRoutes(r)
	initTransactionRoutes(r)
	initQRCodeRoutes(r)
	initBlockchainRoutes(r)
	initMiscRoutes(r)
}

func initUserRoutes(r *gin.RouterGroup) {
	router := r.Group("user/")

	router.GET("info", user.Info)
	router.GET("list", user.List)
	router.POST("register", user.Register)
	router.POST("login", user.Login)
	router.GET("check", user.Check)
	router.GET("access/grant", user.GrantAccessRight)
	router.GET("access/revoke", user.RevokeAccessRight)
}

func initAdminRoutes(r *gin.RouterGroup) {
	router := r.Group("admin/")

	router.GET("info")
}

func initProductRoutes(r *gin.RouterGroup) {
	router := r.Group("product/")

	router.GET("info", product.GetInformation)
	router.GET("list", product.List)
	router.POST("add", product.AddProduct)
}

func initTransactionRoutes(r *gin.RouterGroup) {
	router := r.Group("transaction/")

	router.GET("history", transaction.GetHistory)
	router.GET("last", transaction.GetLastRecord)
	router.POST("add", transaction.AddTransaction)
}

func initQRCodeRoutes(r *gin.RouterGroup) {
	router := r.Group("qrcode/")

	router.GET("generate", qrcode.Generate)
	router.GET("retrieve", qrcode.Retrieve)
}

func initBlockchainRoutes(r *gin.RouterGroup) {
	router := r.Group("blockchain/")

	router.GET("chaincode/list", blockchain.ListChainCode)
}

func initMiscRoutes(r *gin.RouterGroup) {
	r.GET("ping", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.JSON(http.StatusOK, gin.H{
			"cartoon": "Tom and Jerry",
			"names":   []string{"Tom", "Jerry", "Sam"},
		})
	})

	r.Static("assets/", "./assets")
}
