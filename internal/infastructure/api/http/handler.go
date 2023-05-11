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
	mux.HandleFunc("/get", h.getAllTasks)
	mux.HandleFunc("/getById", h.getTaskById)
	mux.HandleFunc("/update", h.updateTask)
	mux.HandleFunc("/delete", h.deleteTask)
	return mux
}
