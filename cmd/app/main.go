package main

import (
	"context"
	"fmt"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/registry"

	"github.com/pyropy/registry/api"
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

	reg, err := registry.NewRegistry(cfg, backend)
	if err != nil {
		panic(fmt.Sprintf("cannot create registry: %s", err))
	}

	fileHandler := api.NewFileRouteHandler(reg)
	server, err := api.NewServer(cfg, fileHandler)
	if err != nil {
		panic(fmt.Sprintf("cannot create server: %s", err))
	}

	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start app: %s", err))
	}
}
