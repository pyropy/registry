package main

import (
	"context"
	"fmt"
	"github.com/pyropy/eth"
	"github.com/pyropy/eth/currency"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/contract"
	"math/big"
)

// TODO: Add logging
func main() {
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		panic(fmt.Sprintf("cannot get config from env: %s", err))
	}

	if cfg.EthPrivateKey == nil {
		panic("eth private key is required")
	}

	backend, err := eth.CreateDialedBackend(context.Background(), cfg.EthRpcUrl)
	if err != nil {
		panic(fmt.Sprintf("cannot create backend: %s", err))
	}

	client, err := eth.NewClient(backend, cfg.EthPrivateKey)
	if err != nil {
		panic(fmt.Sprintf("cannot create eth client: %s", err))
	}

	const gasLimit = 1600000
	valueGwei := big.NewFloat(0.0)
	gasPrice := currency.GWei2Wei(big.NewFloat(39.576))
	txOpts, err := client.NewTransactOpts(context.Background(), gasLimit, gasPrice, valueGwei)
	if err != nil {
		panic(fmt.Sprintf("unable to create transaction opts for deploy: %s", err))
	}

	address, tx, _, err := contract.DeployFileRegistry(txOpts, client.Backend)
	if err != nil {
		panic(fmt.Sprintf("unable to deploy Registry: %s", err))
	}

	if _, err := client.WaitMined(context.Background(), tx); err != nil {
		panic(fmt.Sprintf("waiting for deploy: %s", err))
	}

	fmt.Printf("deployed contract at address: %s\n", address)
}
