package handlers

import "CloudStorage/internal/services"

type Handler struct {
	svc services.ServiceInterface
}

func NewHandler(svc services.ServiceInterface) *Handler {
	return &Handler{
		svc: svc,
	}
}
