package http

import (
	"encoding/json"
	"github.com/romandnk/crud/internal/entities/task"
	"net/http"
)

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var input task.Task

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.Task.Create(ctx, input) // in Create must be context
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]interface{}{
		"id": id,
	})
	w.Write(resp)
	return
}

func (h *Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {

}
