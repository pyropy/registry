package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pyropy/eth"
	"os"
	"strconv"
)

const (
	// Environment variables
	EnvPort                        = "PORT"
	EnvIpfsPort                    = "IPFS_PORT"
	EnvIpfsHost                    = "IPFS_HOST"
	EnvIpfsAutoPin                 = "IPFS_AUTO_PIN"
	EnvEthRpcUrl                   = "ETH_RPC_URL"
	EnvEthPKPath                   = "ETH_PK_PATH"
	EnvEthPKPass                   = "ETH_PK_PASS"
	EnvFileRegistryContractAddress = "FILE_REGISTRY_CONTRACT_ADDRESS"

	// Default fallback values
	DefaultPort                        = "8000"
	DefaultIpfsPort                    = "5001"
	DefaultIpfsHost                    = "localhost"
	DefaultIpfsAutoPin                 = "true"
	DefaultEthRpcUrl                   = "http://localhost:8545"
	DefaultFileRegistryContractAddress = "0x"
)

type Config struct {
	// HTTP Server
	Port int

	// IPFS
	IpfsPort    string
	IpfsHost    string
	IpfsAutoPin bool

	// Ethereum
	EthRpcUrl     string
	EthPrivateKey *ecdsa.PrivateKey

	// File Registry
	FileRegistryContractAddress common.Address
}

var DefaultConfig = &Config{
	Port:        8000,
	IpfsPort:    "5001",
	IpfsHost:    "localhost",
	IpfsAutoPin: true,
	EthRpcUrl:   "http://localhost:8545",
}

func LoadConfigFromEnv() (*Config, error) {
	port, err := strconv.Atoi(getENV(EnvPort, DefaultPort))
	if err != nil {
		return nil, err
	}

	pin, err := strconv.ParseBool(getENV(EnvIpfsAutoPin, DefaultIpfsAutoPin))
	if err != nil {
		return nil, err
	}

	var pk *ecdsa.PrivateKey
	if file := os.Getenv(EnvEthPKPath); file != "" {
		pass := os.Getenv(EnvEthPKPass)
		pk, err = eth.PrivateKeyByKeyFile(file, pass)
		if err != nil {
			return nil, err
		}
	}

	contractAddr := common.HexToAddress(getENV(EnvFileRegistryContractAddress, DefaultFileRegistryContractAddress))
	return &Config{
		Port:                        port,
		IpfsPort:                    getENV(EnvIpfsPort, DefaultIpfsPort),
		IpfsHost:                    getENV(EnvIpfsHost, DefaultIpfsHost),
		IpfsAutoPin:                 pin,
		EthRpcUrl:                   getENV(EnvEthRpcUrl, DefaultEthRpcUrl),
		EthPrivateKey:               pk,
		FileRegistryContractAddress: contractAddr,
	}, nil
}

func getENV(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
