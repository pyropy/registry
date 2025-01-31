package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pyropy/registry/config"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type RouteHandler interface {
	RegisterRoutes(r *gin.RouterGroup)
}

func NewServer(cfg *config.Config, handlers ...RouteHandler) (*http.Server, error) {
	r := gin.Default()
	apiV1 := r.Group("/v1")

	for _, h := range handlers {
		h.RegisterRoutes(apiV1)
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
