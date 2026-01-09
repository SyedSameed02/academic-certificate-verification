package blockchain

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EthClient *ethclient.Client

	CertificateRegistryInstance *CertificateRegistry
	DIDRegistryInstance         *DIDRegistry

	initOnce sync.Once
	initErr  error
)

func Init(
	rpcURL string,
	certificateRegistryAddr string,
	didRegistryAddr string,
) error {

	initOnce.Do(func() {
		EthClient, initErr = ethclient.DialContext(context.Background(), rpcURL)
		if initErr != nil {
			log.Println("❌ Failed to connect to Ethereum:", initErr)
			return
		}
		log.Println("✅ Ethereum client connected")

		CertificateRegistryInstance, initErr = NewCertificateRegistry(
			common.HexToAddress(certificateRegistryAddr),
			EthClient,
		)
		if initErr != nil {
			return
		}

		DIDRegistryInstance, initErr = NewDIDRegistry(
			common.HexToAddress(didRegistryAddr),
			EthClient,
		)
		if initErr != nil {
			return
		}

		log.Println("✅ Smart contract instances initialized")
	})

	if EthClient == nil {
		return errors.New("ethereum client initialization failed")
	}

	return initErr
}
