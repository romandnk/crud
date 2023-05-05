package http

import (
	"github.com/romandnk/crud/internal/infastructure/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", h.createTask)
	mux.HandleFunc("/get", h.createTask)
	mux.HandleFunc("/getById", h.createTask)
	mux.HandleFunc("/update", h.createTask)
	mux.HandleFunc("/delete", h.createTask)
	return mux
}
