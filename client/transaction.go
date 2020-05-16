package client

import (
	"encoding/json"
	"os"

	"github.com/BalusChen/IKHNAIE/types"
)

var (
	ikhnaieOrg1Client *Client
)

func init() {
	ccikhnaieConfig := CCConfig{
		CCID:     "ccikhnaie",
		CCPath:   "github.com/hyperledger/fabric-samples/chaincode/ikhnaie/",
		CCGoPath: os.Getenv("GOPATH"),
	}
	ikhnaieOrg1Client = New(org1CfgPath, "Org1", "Admin", "User1", ccikhnaieConfig)
}

func AddTransaction(foodId string, transaction types.Transaction) error {
	action := "addTransaction"
	transactionBytes, _ := json.Marshal(transaction)

	_, err := ikhnaieOrg1Client.Invoke(peer0Org1, [][]byte{[]byte(action), []byte(foodId), transactionBytes})
	if err != nil {
		return err
	}
	return nil
}

func GetTransactionHistory(foodId string) ([]types.Transaction, error) {
	action := "getTransactionHistory"
	payload, err := ikhnaieOrg1Client.Invoke(peer0Org1, [][]byte{[]byte(action), []byte(foodId)})

	var transactions []types.Transaction
	err = json.Unmarshal(payload, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
