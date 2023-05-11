package http

import (
	"encoding/json"
	"github.com/romandnk/crud/internal/entities"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task entities.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			logrus.Error(err)
			http.Error(w, "unable to unmarshal JSON", http.StatusBadRequest)
			return
		}
		if len(task.Message) == 0 {
			http.Error(w, "message must be not empty", http.StatusBadRequest)
			return
		}

		id, err := h.services.Tasker.Create(r.Context(), task)
		if err != nil || id < 0 {
			logrus.Error(err)
			http.Error(w, "error creating a task", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]int{
			"id": id,
		}); err != nil {
			logrus.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasksList, err := h.services.Tasker.GetAll(r.Context())
		if err != nil {
			logrus.Error(err)
			http.Error(w, "error receiving all tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(tasksList)
		if err != nil {
			logrus.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) getTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			logrus.Error(err)
			http.Error(w, "id must be a number", http.StatusBadRequest)
			return
		}

		task, err := h.services.Tasker.GetById(r.Context(), id)
		if err != nil {
			logrus.Error(err)
			http.Error(w, "error receiving a task", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(task)
		if err != nil {
			logrus.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			logrus.Error(err)
			http.Error(w, "id must be a number", http.StatusBadRequest)
			return
		}

		var newTask entities.Task
		err = json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			logrus.Error(err)
			http.Error(w, "unable to unmarshal JSON", http.StatusBadRequest)
			return
		}
		if len(newTask.Message) == 0 {
			http.Error(w, "message must be not empty", http.StatusBadRequest)
			return
		}

		task, err := h.services.Tasker.Update(r.Context(), id, newTask)
		if err != nil {
			logrus.Error(err)
			http.Error(w, "error updating a task", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(task)
		if err != nil {
			logrus.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			logrus.Error(err)
			http.Error(w, "id must be a number", http.StatusBadRequest)
			return
		}

		err = h.services.Tasker.Delete(r.Context(), id)
		if err != nil {
			logrus.Error(err)
			http.Error(w, "error deleting a task", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
