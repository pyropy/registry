package main

import (
	"fmt"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/registry"

	"github.com/pyropy/registry/api"
)

func main() {
	reg, err := registry.NewRegistry(&config.DefaultConfig)
	if err != nil {
		panic(fmt.Sprintf("cannot create registry: %s", err))
	}

	fileHandler := api.NewFileRouteHandler(reg)
	server, err := api.NewServer(&config.DefaultConfig, fileHandler)
	if err != nil {
		panic(fmt.Sprintf("cannot create server: %s", err))
	}

	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start app: %s", err))
	}

}
