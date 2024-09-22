package handlers

import (
	"TimeBookingAPI/internal/repository"
	"TimeBookingAPI/internal/service"
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	answer := service.Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.GetJson())
	}()

	d, err := io.ReadAll(r.Body)
	if err != nil || len(d) == 0 {
		h.LogError(r, err)
		answer.Status = http.StatusBadRequest
		return
	}

	newUser := repository.NewUserModel()
	err = json.Unmarshal(d, &newUser)
	if err != nil {
		h.LogError(r, err)
		answer.Status = http.StatusBadRequest
		return
	}

	answer, err = h.service.CreateUser(r.Context(), newUser)
	if err != nil {
		h.LogError(r, err)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	answer := service.Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.GetJson())
	}()

	username := r.URL.Query().Get("username")

	answer, err := h.service.DeleteUser(r.Context(), username)
	if err != nil {
		h.LogError(r, err)
		answer.Status = http.StatusInternalServerError
		return
	}
}
