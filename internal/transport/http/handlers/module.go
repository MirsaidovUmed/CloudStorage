package handlers

import (
	"CloudStorage/internal/services"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	svc    services.ServiceInterface
	logger *logrus.Logger
}

func NewHandler(svc services.ServiceInterface, logger *logrus.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}
