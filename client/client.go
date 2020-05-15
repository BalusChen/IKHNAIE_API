package client

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	chaincodeName    = "mycc"
	chaincodeVersion = "0"
	channelName      = "myc"
)

type ClientConfig struct {
}

type Client struct {
	// Fabric network information
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgUser    string

	// sdk clients
	SDK *fabsdk.FabricSDK
	rc  *resmgmt.Client
	cc  *channel.Client

	// Same for each peer
	ChannelID string
	CCID      string // chaincode ID, eq name
	CCPath    string // chaincode source path, 是GOPATH下的某个目录
	CCGoPath  string // GOPATH used for chaincode
}

func New(cfgPath, org, admin, user string) *Client {
	client := &Client{
		ConfigPath: cfgPath,
		OrgName:    org,
		OrgAdmin:   admin,
		OrgUser:    user,

		CCID:      "example2",
		CCPath:    "github.com/hyperledger/fabric-samples/chaincode/chaincode_example02/go/",
		CCGoPath:  os.Getenv("GOPATH"),
		ChannelID: "mychannel",
	}

	sdk, err := fabsdk.New(config.FromFile(client.ConfigPath))
	if err != nil {
		log.Panicf("Failed to create fabric sdk, err: %v", err)
	}

	client.SDK = sdk

	rcp := sdk.Context(fabsdk.WithUser(client.OrgAdmin), fabsdk.WithOrg(client.OrgName))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		log.Panicf("Failed to create resource client")
	}

	log.Println("Succeed to initialize resource client")

	ccp := sdk.ChannelContext(client.ChannelID, fabsdk.WithUser(client.OrgUser))
	cc, err := channel.New(ccp)
	if err != nil {
		log.Panicf("Failed to create channel client")
	}

	log.Println("Succeed to initialize channel client")

	client.rc = rc
	client.cc = cc
	return client
}

func (client *Client) Query(peer, key string) error {
	rawReq := channel.Request{
		ChaincodeID: client.CCID,
		Fcn:         "query",
		Args:        [][]byte{[]byte(key)},
	}
	req := channel.WithTargetEndpoints(peer)
	resp, err := client.cc.Query(rawReq, req)
	if err != nil {
		return errors.New(fmt.Sprintf("query chaincode failed, err: %v", err))
	}

	log.Printf("%+v", resp)
	return nil
}
