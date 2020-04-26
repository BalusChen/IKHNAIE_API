package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInformation(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello product")
}
