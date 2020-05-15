package client

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

const (
	chaincodeName    = "ccikhnaie"
	chaincodeVersion = "1.0"
	channelName      = "mychannel"
)

type CCConfig struct {
	CCID      string // chaincode ID, eq name
	CCVersion string // chaincode version
	CCPath    string // chaincode source path, 是 GOPATH 下的某个目录
	CCGoPath  string // GOPATH used for chaincode
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
	CCCfg CCConfig
}

func New(cfgPath, org, admin, user string, ccCfg CCConfig) *Client {
	client := &Client{
		ConfigPath: cfgPath,
		OrgName:    org,
		OrgAdmin:   admin,
		OrgUser:    user,

		ChannelID: "mychannel",
		CCCfg: ccCfg,
	}

	sdk, err := fabsdk.New(config.FromFile(client.ConfigPath))
	if err != nil {
		log.Panicf("[ClientNew] failed to create fabric sdk, err: %v", err)
	}

	client.SDK = sdk

	rcp := sdk.Context(fabsdk.WithUser(client.OrgAdmin), fabsdk.WithOrg(client.OrgName))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		log.Panicf("[ClientNew] failed to create resource client")
	}

	log.Println("[ClientNew] succeed to initialize resource client")

	ccp := sdk.ChannelContext(client.ChannelID, fabsdk.WithUser(client.OrgUser))
	cc, err := channel.New(ccp)
	if err != nil {
		log.Panicf("[ClientNew] failed to create channel client")
	}

	log.Println("[ClientNew] succeed to initialize channel client")

	client.rc = rc
	client.cc = cc
	return client
}
