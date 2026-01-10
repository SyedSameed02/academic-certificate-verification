package blockchain

import (
	"log"
	"os"

	"backend-go/contracts"
	"context"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
)

type CertificateService struct {
	contract *contracts.CertificateRegistry
	client   *Client
}

func NewCertificateService(client *Client) *CertificateService {
	addr := common.HexToAddress(os.Getenv("CERTIFICATE_REGISTRY_ADDRESS"))

	contract, err := contracts.NewCertificateRegistry(addr, client.Eth)
	if err != nil {
		log.Fatal("failed to init CertificateRegistry:", err)
	}

	return &CertificateService{
		contract: contract,
		client:   client,
	}
}

func (c *CertificateService) IssueCertificate(certHash [32]byte) error {
	tx, err := c.contract.IssueCertificate(c.client.Auth, certHash)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(context.Background(), c.client.Eth, tx)

	return err
}

func (c *CertificateService) RevokeCertificate(certHash [32]byte) error {
	tx, err := c.contract.RevokeCertificate(c.client.Auth, certHash)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(context.Background(), c.client.Eth, tx)

	return err
}

func (c *CertificateService) GetCertificate(certHash [32]byte) (issuer common.Address, revoked bool, issuedAt uint64, err error) {
	cert, err := c.contract.GetCertificate(c.client.CallOpts, certHash)
	if err != nil {
    	return common.Address{}, false, 0, err
	}

	return cert.Issuer, cert.Revoked, cert.IssuedAt.Uint64(), nil

}

func (c *CertificateService) Exists(certHash [32]byte) (bool, error) {
	return c.contract.Exists(c.client.CallOpts, certHash)
}
