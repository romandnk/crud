package http

import (
	"net/http"
)

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	id, err := h.services.Create() // in Create must be context

}

func (h *Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {

}
