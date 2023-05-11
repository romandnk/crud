package http

import (
	"encoding/json"
	"github.com/romandnk/crud/internal/entities"
	"net/http"
	"strconv"
)

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task entities.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "unable to unmarshal JSON", err)
			return
		}
		if len(task.Message) == 0 {
			newErrorResponse(w, http.StatusBadRequest, "message must be not empty", err)
			return
		}

		id, err := h.services.Tasker.Create(r.Context(), task)
		if err != nil || id < 0 {
			newErrorResponse(w, http.StatusInternalServerError, "error creating a task", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]int{
			"id": id,
		}); err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
			return
		}
		return
	} else {
		newErrorResponse(w, http.StatusMethodNotAllowed, "only post method is available", nil)
		return
	}
}

func (h *Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasksList, err := h.services.Tasker.GetAll(r.Context())
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, "error receiving all tasks", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(tasksList)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
			return
		}
		return
	} else {
		newErrorResponse(w, http.StatusMethodNotAllowed, "only get method is available", nil)
		return
	}
}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "wrong parameters", err)
			return
		}

		task, err := h.services.Tasker.GetById(r.Context(), id)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, "error receiving a task", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(task)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
			return
		}
		return
	} else {
		newErrorResponse(w, http.StatusMethodNotAllowed, "only get method is available", nil)
		return
	}
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "wrong parameters", err)
			return
		}

		var newTask entities.Task
		err = json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "unable to unmarshal JSON", err)
			return
		}
		if len(newTask.Message) == 0 {
			newErrorResponse(w, http.StatusBadRequest, "message must be not empty", err)
			return
		}

		task, err := h.services.Tasker.Update(r.Context(), id, newTask)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, "error updating a task", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(task)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error(), err)
			return
		}
		return
	} else {
		newErrorResponse(w, http.StatusMethodNotAllowed, "only put method is available", nil)
		return
	}
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, "wrong parameters", err)
			return
		}

		err = h.services.Tasker.Delete(r.Context(), id)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, "error deleting a task", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		newErrorResponse(w, http.StatusMethodNotAllowed, "only delete method is available", nil)
		return
	}
}
