package main

import (
	"context"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/contract"
	"github.com/rs/zerolog/log"
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

	address, _, err := contract.DeployContract(ctx, client, DefaultDeployGasLimit)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to deploy contract")
	}

	log.Info().
		Str("address", address.String()).
		Msg("deployed contract")

}
