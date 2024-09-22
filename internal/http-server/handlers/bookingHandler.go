package handlers

import (
	"TimeBookingAPI/internal/repository"
	"TimeBookingAPI/internal/service"
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
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

	newBooking := repository.NewBookingModel()
	err = json.Unmarshal(d, &newBooking)
	if err != nil {
		h.LogError(r, err)
		answer.Status = http.StatusBadRequest
		return
	}

	answer, err = h.service.CreateBooking(r.Context(), newBooking)
	if err != nil {
		h.LogError(r, err)
		return
	}
}

func (h *Handler) FindBooking(w http.ResponseWriter, r *http.Request) {
	answer := service.Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.GetJson())
	}()

	username := r.URL.Query().Get("username")
	answer, err := h.service.FindBooking(r.Context(), username)
	if err != nil {
		h.LogError(r, err)
		return
	}
}

func (h *Handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	answer := service.Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.GetJson())
	}()

	id := r.URL.Query().Get("id")
	answer, err := h.service.DeleteBooking(r.Context(), id)
	if err != nil {
		h.LogError(r, err)
		return
	}
}
