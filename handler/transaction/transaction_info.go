package transaction

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BalusChen/IKHNAIE/types"
	"github.com/BalusChen/IKHNAIE_API/client"
	"github.com/BalusChen/IKHNAIE_API/constant"
	"github.com/gin-gonic/gin"
)

type transactionInfo struct {
	FoodID     string  `form:"food_id" binding:"required"`
	SellerName string  `form:"seller_name" binding:"required"`
	SellerID   string  `form:"seller_id" binding:"required"`
	BuyerName  string  `form:"buyer_name" binding:"required"`
	BuyerID    string  `form:"buyer_id" binding:"required"`
	TradePlace string  `form:"trade_place" binding:"required"`
	Number     int64   `form:"number" binding:"required"`
	Price      float64 `form:"price" binding:"required"`
}

func AddTransaction(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var transaction transactionInfo
	if err := ctx.ShouldBind(&transaction); err != nil {
		log.Printf("[AddTransaction] invalid params: %+v", ctx.Params)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	tx := types.Transaction{
		TradeTime:  time.Now(),
		TradePlace: transaction.TradePlace,
		SellerName: transaction.SellerName,
		SellerID:   transaction.SellerID,
		BuyerName:  transaction.BuyerName,
		BuyerID:    transaction.BuyerID,
		Number:     transaction.Number,
		Price:      transaction.Price,
	}

	fmt.Printf("tx:\n%+v\n", tx)

	err := client.AddTransaction(transaction.FoodID, tx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": constant.StatusCode_CallBlockChainError,
			"status_msg":  constant.StatusMsg_CallBalockChainError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":    constant.StatusCode_OK,
		"status_message": constant.StatusMsg_OK,
	})
}

func GetHistory(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	foodID, found := ctx.GetQuery("food_id")
	if !found {
		log.Printf("[GetHistory] invalid params, \"food_id\" is not specified, params: %v\n", ctx.Params)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusMsg_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	log.Printf("[GetHistory] query transaction history for foodID=%s\n", foodID)

	transactionHistory, err := client.GetTransactionHistory(foodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": constant.StatusCode_CallBlockChainError,
			"status_msg":  constant.StatusMsg_CallBalockChainError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":         constant.StatusCode_OK,
		"status_msg":          constant.StatusMsg_OK,
		"transaction_history": transactionHistory,
	})
}

func GetLastRecord(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":    constant.StatusCode_MethodONotImplemented,
		"status_message": constant.StatusMsg_MethodNotImplemented,
	})
}
