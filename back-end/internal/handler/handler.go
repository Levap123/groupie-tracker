package handler

import (
	"groupie-tracker/back-end/internal/service"
	"groupie-tracker/back-end/pkg/logger"
	"net/http"
)

type Handler struct {
	service *service.Service
	logger  *logger.Logger
}

func NewHandler(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	router := http.NewServeMux()
	router.Handle("/api/artists/", CORS(http.HandlerFunc(h.Artist)))
	router.Handle("/api/concerts/", CORS(http.HandlerFunc(h.Concert)))
	return router
}
