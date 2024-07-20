package middleware

import (
	"CloudStorage/internal/services"
	"CloudStorage/pkg/config"
	"net/http"
)

type Middleware struct {
	config  *config.Config
	service services.ServiceInterface
}

type MiddlewareInterface interface {
	TimeDuration(next http.Handler) http.Handler
	JWT(next http.Handler) http.Handler
}

func NewMiddleware(config *config.Config, svc services.ServiceInterface) MiddlewareInterface {
	return &Middleware{
		config:  config,
		service: svc,
	}
}
