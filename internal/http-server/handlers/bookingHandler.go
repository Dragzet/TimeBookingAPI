package handlers

import (
	"TimeBookingAPI/internal/bookingModule"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.getJson())
	}()

	d, err := io.ReadAll(r.Body)
	if err != nil || len(d) == 0 {
		h.LogError(r, err)
		answer.Status = http.StatusBadRequest
		return
	}

	newBooking := bookingModule.New()
	err = json.Unmarshal(d, &newBooking)
	if err != nil || newBooking.Delta < 1 {
		h.LogError(r, err)
		answer.Status = http.StatusBadRequest
		return
	}

	newBooking.EndTime = time.Now().Add(time.Hour * time.Duration(newBooking.Delta))
	err = h.bookingStorage.Create(r.Context(), newBooking)
	if err != nil {
		h.LogError(r, err)
		answer.Status = http.StatusInternalServerError
		return
	}
}

func (h *Handler) FindBooking(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.getJson())
	}()

	username := r.URL.Query().Get("username")
	if username == "" {
		answer.Status = http.StatusBadRequest
		return
	}

	bookings, err := h.bookingStorage.FindAll(r.Context(), username)
	if err != nil {
		h.LogError(r, err)
		answer.Status = http.StatusInternalServerError
		return
	}

	answer.Data = bookings
}

func (h *Handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.getJson())
	}()

	id := r.URL.Query().Get("id")
	if id == "" {
		h.LogError(r, fmt.Errorf("no id provided"))
		answer.Status = http.StatusBadRequest
		return
	}

	err := h.bookingStorage.Delete(r.Context(), id)
	if err != nil {
		h.LogError(r, err)
		answer.Status = http.StatusInternalServerError
		return
	}
}
