package client

import (
	"fmt"
	"os"
)

func helloCC() {
	cchelloConfig := CCConfig{
		CCID:     "exacc",
		CCPath:   "github.com/chaincode/chaincode_example02/go/",
		CCGoPath: os.Getenv("GOPATH"),
	}
	org1Client := New(org1CfgPath, "Org1", "Admin", "User1", cchelloConfig)
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
