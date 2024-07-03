package main

import (
	"context"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
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

	backend, err := eth.CreateDialedBackend(ctx, cfg.EthRpcUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create eth backend")
	}

	reg, err := registry.NewRegistry(cfg, backend)
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
