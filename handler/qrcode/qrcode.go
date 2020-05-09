package qrcode

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

const (
	qrCodeURLTpl = "127.0.0.1:9877/ikhnaie/v1/qrcode/retrieve?food_id=%d"
	qrCodePngTpl = "assets/images/qrcode/%d.png"
)

func Generate(ctx *gin.Context) {
	foodIDStr, ok := ctx.GetQuery("food_id")
	if !ok {
		/* TODO: log */
		return
	}

	foodID, err := strconv.ParseInt(foodIDStr, 10, 64)
	if err != nil {
		/* TODO: log */
		return
	}

	url := fmt.Sprintf(qrCodeURLTpl, foodID)
	err = qrcode.WriteFile(url, qrcode.Medium, 256, fmt.Sprintf(qrCodePngTpl, foodID))
	if err != nil {
		/* TODO: log */
		return
	}
}

func Retrieve(ctx *gin.Context) {
	foodIdStr, ok := ctx.GetQuery("food_id")
	if !ok {
		ctx.String(http.StatusBadRequest, "no food_id param found")
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("query food_id=%s", foodIdStr))

	/* TODO: rpc blockchain */
}
