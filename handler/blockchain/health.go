package blockchain

import (
	"net/http"

	"github.com/BalusChen/IKHNAIE_API/constant"
	"github.com/gin-gonic/gin"
)

type chainCodeInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
	ID      string `json:"id"`
}

func ListChainCode(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ccs := []chainCodeInfo{
		{
			"cchello1",
			"1.0",
			"github.com/chaincode/hello/",
			"a2b44a68e6b122fabccd8f9d024227ed8144737aff3fe8f0cbc324160d9fe5d6",
		},
		{
			"ccikhnaie",
			"1.0",
			"github.com/chaincode/ikhnaie",
			"1a6904df24419c7ad1e604ab7920e55c002656a52f101714de1a1b6c988df1e4",
		},
		{
			"ccikhnaie",
			"2.0",
			"github.com/chaincode/ikhnaie",
			"f99eb1d824b5bfa116f7741cdab76393cf7027df5ab2cfbd5e6833418e87853a",
		},
		{
			"exacc",
			"1.0",
			"github.com/chaincode/chaincode_example02/go/",
			"e3acd97bb4d9dbd868c8e97b3d7d7802163f6ba886859d7af53be40b38d42e6d",
		},
	}
	ctx.JSON(http.StatusOK, gin.H{
		"chaincodes":  ccs,
		"status_code": http.StatusOK,
		"status_msg":  constant.StatusMsg_OK,
	})
}
