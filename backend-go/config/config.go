package config

type Config struct {
	Server struct {
		Port int
	}
	BlockChain struct {
		RPCUrl              string
		CertificateRegistry string
		DIDRegistry         string
	}
}
