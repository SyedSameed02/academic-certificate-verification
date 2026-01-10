package blockchain

import (
	"log"
	"os"

	"backend-go/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
)

type DIDService struct {
	contract *contracts.DIDRegistry
}

func NewDIDService(client *Client) *DIDService {
	addr := common.HexToAddress(os.Getenv("DID_REGISTRY_ADDRESS"))

	contract, err := contracts.NewDIDRegistry(addr, client.Eth)
	if err != nil {
		log.Fatal("failed to init DIDRegistry:", err)
	}

	return &DIDService{contract: contract}
}

func (d *DIDService) IsValidIssuer(addr common.Address) (bool, error) {
	return d.contract.IsValidIssuer(clientCall(), addr)
}
func clientCall() *bind.CallOpts {
	return &bind.CallOpts{}
}
