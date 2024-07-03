package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kelseyhightower/envconfig"
	"github.com/pyropy/eth"
	"os"
)

type RawConfig struct {
	HttpPort                    int    `envconfig:"HTTP_PORT" default:"8000"`
	IpfsUrl                     string `envconfig:"IPFS_URL" default:"http://localhost:5001"`
	IpfsAutoPin                 bool   `envconfig:"IPFS_AUTO_PIN" default:"true"`
	EthRpcUrl                   string `envconfig:"ETH_RPC_URL" default:"http://localhost:8545"`
	EthGasLimit                 uint64 `envconfig:"ETH_GAS_LIMIT" default:"21000"`
	EthPrivateKeyFile           string `envconfig:"ETH_PRIVATE_KEY_FILE" default:""`
	EthPrivateKeyPasswordFile   string `envconfig:"ETH_PRIVATE_KEY_PASSWORD_FILE" default:""`
	FileRegistryContractAddress string `envconfig:"FILE_REGISTRY_CONTRACT_ADDRESS" default:"0x"`
}

type Config struct {
	// HTTP Server
	Port int

	// IPFS
	IpfsUrl     string
	IpfsAutoPin bool

	// Ethereum
	EthRpcUrl     string
	EthGasLimit   uint64
	EthPrivateKey *ecdsa.PrivateKey

	// File Registry
	FileRegistryContractAddress common.Address
}

func LoadConfigFromEnv() (*Config, error) {
	var raw RawConfig
	err := envconfig.Process("", &raw)
	if err != nil {
		return nil, err
	}

	pk, err := readEthPrivateKey(raw)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:                        raw.HttpPort,
		IpfsUrl:                     raw.IpfsUrl,
		EthRpcUrl:                   raw.EthRpcUrl,
		IpfsAutoPin:                 raw.IpfsAutoPin,
		EthPrivateKey:               pk,
		EthGasLimit:                 raw.EthGasLimit,
		FileRegistryContractAddress: common.HexToAddress(raw.FileRegistryContractAddress),
	}, nil
}

func readEthPrivateKey(raw RawConfig) (pk *ecdsa.PrivateKey, err error) {
	var password string
	if raw.EthPrivateKeyPasswordFile != "" {
		p, err := os.ReadFile(raw.EthPrivateKeyPasswordFile)
		if err != nil {
			return nil, err
		}

		password = string(p)
	}

	if raw.EthPrivateKeyFile != "" {
		pk, err = eth.PrivateKeyByKeyFile(raw.EthPrivateKeyFile, password)
		if err != nil {
			return nil, err
		}
	}

	return pk, nil
}
