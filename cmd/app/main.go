package main

import (
	"context"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/ipfs"
	"github.com/pyropy/registry/registry"
	"github.com/rs/zerolog/log"

	"github.com/pyropy/registry/api"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config from env")
	}

	if cfg.EthPrivateKey == nil {
		log.Fatal().Msg("private key is required")
	}

	ipfsClient, err := ipfs.NewDialedClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create ipfs client")
	}

	backend, err := eth.CreateDialedBackend(ctx, cfg.EthRpcUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create eth backend")
	}

	defer backend.Close()

	reg, err := registry.NewRegistry(cfg, ipfsClient, backend)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create registry servic")
	}

	fileHandler := api.NewFileRouteHandler(reg)
	server, err := api.NewServer(cfg, fileHandler)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create registry server")
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("failed to start the app")
	}
}
