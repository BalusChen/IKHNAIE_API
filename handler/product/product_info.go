package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInformation(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: get product information")
}

func AddProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: add new product")
}
