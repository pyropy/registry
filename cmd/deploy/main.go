package main

import (
	"context"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/contract"
	"github.com/rs/zerolog/log"
	"math/big"
)

const DefaultDeployGasLimit = 6_000_000

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config from env")
	}

	if cfg.EthPrivateKey == nil {
		log.Fatal().Msg("private key is required")
	}

	backend, err := eth.CreateDialedBackend(ctx, cfg.EthRpcUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create eth backend")
	}

	client, err := eth.NewClient(backend, cfg.EthPrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create eth client")
	}

	valueGwei := big.NewFloat(0.0)
	txOpts, err := client.NewTransactOpts(context.Background(), DefaultDeployGasLimit, nil, valueGwei)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create transaction opts for deploy")
	}

	address, tx, _, err := contract.DeployFileRegistry(txOpts, client.Backend)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to deploy Registry")
	}

	if _, err := client.WaitMined(context.Background(), tx); err != nil {
		log.Fatal().Err(err).Msg("failed to get deployment tx")
	}

	log.Info().
		Str("address", address.String()).
		Msg("deployed contract")

}
