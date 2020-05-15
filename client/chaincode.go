package client

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
)

func (client *Client) InstallCC(version string, peer string) error {
	targetPeer := resmgmt.WithTargetEndpoints(peer)

	ccPkg, err := gopackager.NewCCPackage(client.CCPath, client.CCGoPath)
	if err != nil {
		return errors.New(fmt.Sprintf("pack chaincode failed, err: %v", err))
	}

	req := resmgmt.InstallCCRequest{
		Name:    client.CCID,
		Path:    client.CCPath,
		Version: version,
		Package: ccPkg,
	}
	resp, err := client.rc.InstallCC(req, targetPeer)
	if err != nil {
		return errors.New(fmt.Sprintf("install chaincode failed, err: %v", err))
	}
	_ = resp

	return nil
}

func (client *Client) InstantiateCC(version string, peer string) (fab.TransactionID, error) {
	return "", nil
}
