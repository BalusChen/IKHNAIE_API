package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/BalusChen/IKHNAIE/types"
	"github.com/BalusChen/IKHNAIE_API/client"
	"github.com/BalusChen/IKHNAIE_API/views"
	"github.com/gin-gonic/gin"
)

func main() {
	ikhnaieCC()
	router := gin.Default()

	views.InitRoutes(router)

	// panic(router.Run(":9877"))
}

const (
	org1CfgPath = "./config/org1sdk-config.yaml"
	org2CfgPath = "./config/org2sdk-config.yaml"

	peer0Org1 = "peer0.org1.example.com"
	peer0Org2 = "peer0.org2.example.com"
)

func helloCC() {
	cchelloConfig := client.CCConfig{
		CCID:     "exacc",
		CCPath:   "github.com/chaincode/chaincode_example02/go/",
		CCGoPath: os.Getenv("GOPATH"),
	}
	org1Client := client.New(org1CfgPath, "Org1", "Admin", "User1", cchelloConfig)
	payload, err := org1Client.Query(peer0Org1, "query", [][]byte{[]byte("a")})
	if err != nil {
		panic(err)
	}
	fmt.Printf("payload: %s", string(payload))

	_, err = org1Client.Invoke(peer0Org1, [][]byte{[]byte("a"), []byte("b"), []byte("200")})
	if err != nil {
		panic(err)
	}
}

func ikhnaieCC() {
	ccikhnaieConfig := client.CCConfig{
		CCID:     "ccikhnaie",
		CCPath:   "github.com/hyperledger/fabric-samples/chaincode/ikhnaie/",
		CCGoPath: os.Getenv("GOPATH"),
	}
	org1Client := client.New(org1CfgPath, "Org1", "Admin", "User1", ccikhnaieConfig)

	action := "addTransaction"
	foodId := "0000001"
	transaction := types.Transaction{
		TradeTime:  time.Now(),
		TradePlace: "New York",
		SellerName: "Jerry",
		SellerID:   "3333333",
		BuyerName:  "Sam",
		BuyerID:    "4444444",
		Number:     24,
		Price:      30,
	}

	transactionBytes, _ := json.Marshal(transaction)
	_, err := org1Client.Invoke(peer0Org1, [][]byte{[]byte(action), []byte(foodId), transactionBytes})
	if err != nil {
		panic(err)
	}

	action = "getTransactionHistory"
	payload, err := org1Client.Invoke(peer0Org1, [][]byte{[]byte(action), []byte(foodId)})
	if err != nil {
		panic(err)
	}

	fmt.Printf("transactionsJSON:\n%v\n", string(payload))

	var transactions []types.Transaction
	err = json.Unmarshal(payload, &transactions)
	if err != nil {
		panic(err)
	}

	fmt.Printf("transactions:\n%v\n", transactions)
}
