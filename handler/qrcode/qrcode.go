package qrcode

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BalusChen/IKHNAIE_API/constant"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

const (
	qrCodeUrlTpl     = "http://localhost:9877/ikhnaie/v1/assets/images/qrcode/%d.png"
	qrCodePngTpl     = "assets/images/qrcode/%d.png"
	qrCodeContentTpl = "http://192.168.43.29/ikhnaie/v1/qrcode/retrieve?food_id=%d"
)

func Generate(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	foodIDStr, ok := ctx.GetQuery("food_id")
	if !ok {
		log.Print("[QRCodeGenerate] no food_id specified")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	foodID, err := strconv.ParseInt(foodIDStr, 10, 64)
	if err != nil {
		log.Printf("[QRCodeGenerate] invalid food_id: %q, err: %v", foodIDStr, err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	url := fmt.Sprintf(qrCodeContentTpl, foodID)
	qrcodeName := fmt.Sprintf(qrCodePngTpl, foodID)
	err = qrcode.WriteFile(url, qrcode.Medium, 256, qrcodeName)
	if err != nil {
		log.Printf("[QRCodeGenerate] write qrcode to file failed, err: %v", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"qrcode_url":  fmt.Sprintf(qrCodeUrlTpl, foodID),
		"status_code": constant.StatusCode_OK,
		"status_msg":  constant.StatusMsg_OK,
	})
}

func Retrieve(ctx *gin.Context) {
	foodIdStr, ok := ctx.GetQuery("food_id")
	if !ok {
		log.Print("[QRCodeRetrieve] bad qrcode: no food_id specified")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("query food_id=%s", foodIdStr))

	/* TODO: rpc blockchain */
}
