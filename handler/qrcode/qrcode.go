package qrcode

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

const (
	qrCodeURLTpl = "192.168.1.5:9877/ikhnaie/v1/qrcode/retrieve?food_id=%d"
	qrCodePngTpl = "assets/images/qrcode/%d.png"
)

func Generate(ctx *gin.Context) {
	foodIDStr, ok := ctx.GetQuery("food_id")
	if !ok {
		log.Print("[GenerateQRCode] no food_id specified")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"reason": "no food_id specified",
		})
		return
	}

	foodID, err := strconv.ParseInt(foodIDStr, 10, 64)
	if err != nil {
		log.Printf("[GenerateQRCode] invalid food_id: %q, err: %v", foodIDStr, err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid food_id",
		})
		return
	}

	url := fmt.Sprintf(qrCodeURLTpl, foodID)
	qrcodeName := fmt.Sprintf(qrCodePngTpl, foodID)
	err = qrcode.WriteFile(url, qrcode.Medium, 256, qrcodeName)
	if err != nil {
		log.Printf("[GenerateQRCode] write qrcode to file failed, err: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"reason": "call qrcode generator failed",
		})
		return
	}

	data, _ := ioutil.ReadFile(qrcodeName)
	_, _ = ctx.Writer.Write(data)
}

func Retrieve(ctx *gin.Context) {
	foodIdStr, ok := ctx.GetQuery("food_id")
	if !ok {
		log.Print("[GenerateQRCode] bad qrcode: no food_id specified")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"reason": "bad qrcode, no food_id specified",
		})
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("query food_id=%s", foodIdStr))

	/* TODO: rpc blockchain */
}
