package client

import (
	"errors"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
)

func (client *Client) InstallCC(version string, peer string) error {
	targetPeer := resmgmt.WithTargetEndpoints(peer)

	ccPkg, err := gopackager.NewCCPackage(client.CCCfg.CCPath, client.CCCfg.CCGoPath)
	if err != nil {
		return errors.New(fmt.Sprintf("pack chaincode failed, err: %v", err))
	}

	req := resmgmt.InstallCCRequest{
		Name:    client.CCCfg.CCID,
		Path:    client.CCCfg.CCPath,
		Version: version,
		Package: ccPkg,
	}
	resp, err := client.rc.InstallCC(req, targetPeer)
	if err != nil {
		return errors.New(fmt.Sprintf("install chaincode failed, err: %v", err))
	}
	// TODO: check other errors
	_ = resp

	log.Printf("[InstallCC] succeed to install chaincode. name: %q, version: %q, peer: %q", client.CCCfg.CCPath, version, peer)

	return nil
}

func (client *Client) InstantiateCC(version string, peer string) (fab.TransactionID, error) {

	log.Printf("[InstantiateCC] succeed to instantiate chaincode. name: %q, version: %q, peer: %q", client.CCCfg.CCPath, version, peer)

	return "", nil
}

func (client *Client) Query(peer string, fcn string, args [][]byte) ([]byte, error) {
	rawReq := channel.Request{
		ChaincodeID: client.CCCfg.CCID,
		Fcn:         fcn,
		Args:        args,
	}
	req := channel.WithTargetEndpoints(peer)
	resp, err := client.cc.Query(rawReq, req)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (client *Client) Invoke(peer string, args [][]byte) ([]byte, error) {
	rawReq := channel.Request{
		ChaincodeID: client.CCCfg.CCID,
		Fcn:         "invoke",
		Args:        args,
	}
	req := channel.WithTargetEndpoints(peer)
	resp, err := client.cc.Execute(rawReq, req)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
