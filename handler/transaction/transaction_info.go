package transaction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	// FoodID  int64   `json:"food_id"` /* key */
	TradeTime  time.Time `json:"trade_time"`  // 交易时间
	TradePlace string    `json:"trade_place"` // 交易地点
	SellerName string    `json:"seller_name"` // 售卖者名字
	SellerID   string    `json:"seller_id"`   // 售卖者身份证号码
	BuyerName  string    `json:"buyer_name"`  // 购买者名字
	BuyerID    string    `json:"buyer_id"`    // 购买者身份证号码
	Number     int64     `json:"number"`      // 交易数目
	Price      float64   `json:"price"`       // 单价
}

type Product struct {
	FoodID       int64     `json:"food_id"`       // 农产品 ID（唯一标识）
	FoodName     string    `json:"food_name"`     // 农产品名
	BirthAddress string    `json:"birth_address"` // 产地
	Birthday     time.Time `json:"birthday"`      // 生产日期
	ShelfLife    int       `json:"shelf_life"`    // 保质期（天）
	// Ingredients []Product     `json:"ingredients"` // 组成材料
}

func GetHistory(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	product := Product{
		FoodID:       1,
		FoodName:     "土豆",
		Birthday:     time.Now(),
		BirthAddress: "BeiJing",
		ShelfLife:    100,
	}
	records := []Transaction{
		{
			time.Now(),
			"BeiJing",
			"Tom",
			"1111111",
			"Jerry",
			"2222222",
			501,
			79,
		},
		{
			time.Now().Add(time.Hour * 12),
			"New York",
			"Jerry",
			"2222222",
			"Sam",
			"3333333",
			300,
			30,
		},
		{
			time.Now().Add(time.Hour * 360),
			"London",
			"Sam",
			"3333333",
			"balus",
			"4444444",
			300,
			90,
		},
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product_info":        product,
		"transaction_history": records,
	})
}

func AddTransaction(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: add transaction")
}
