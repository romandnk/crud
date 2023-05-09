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

	id, err := h.services.Task.Create(r.Context(), input) // in Create must be context
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]interface{}{
		"id": id,
	})
	w.Write(resp)
}

type getAllListsResponse struct {
	Data []task.Task `json:"data"`
}

func (h *Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	lists, err := h.services.Task.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(getAllListsResponse{
		Data: lists,
	})
	w.Write(resp)
}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	getTask, err := h.services.Task.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[int]task.Task{
		id: getTask,
	})
	w.Write(resp)
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	var input task.Task
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateTask, err := h.services.Task.Update(r.Context(), input.Id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[int]task.Task{
		input.Id: updateTask,
	})
	w.Write(resp)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	var input task.Task
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.services.Task.Delete(r.Context(), input.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
