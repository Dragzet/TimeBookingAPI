package handlers

import (
	"TimeBookingAPI/internal/bookingModule"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CreateBooking godoc
// @Summary Create a booking
// @Description Create a new booking with the provided details
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body bookingModule.BookingModel true "Booking data"
// @Success 201 {object} Answer
// @Failure 400 {object} Answer "Invalid input"
// @Failure 500 {object} Answer "Internal server error"
// @Router /booking [post]
func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.getJson())
	}()

	d, err := io.ReadAll(r.Body)
	if err != nil || len(d) == 0 {
		answer.Status = http.StatusBadRequest
		return
	}

	newBooking := bookingModule.New()
	err = json.Unmarshal(d, &newBooking)
	if err != nil {
		answer.Status = http.StatusBadRequest
		return
	}

	newBooking.EndTime = time.Now().Add(time.Hour * time.Duration(newBooking.Delta))
	err = h.bookingStorage.Create(r.Context(), newBooking)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return
	}
}

// FindBooking godoc
// @Summary Find a booking
// @Description Get a booking by its ID
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} Answer "Booking details"
// @Failure 404 {object} Answer "Booking not found"
// @Failure 500 {object} Answer "Internal server error"
// @Router /booking?username={username} [get]
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
		fmt.Print(err.Error())
		answer.Status = http.StatusInternalServerError
		return
	}

	answer.Data = bookings
}

// DeleteBooking godoc
// @Summary Delete a booking
// @Description Delete a booking by its ID
// @Tags booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 204 {object} Answer "No content"
// @Failure 404 {object} Answer "Booking not found"
// @Failure 500 {object} Answer "Internal server error"
// @Router /booking?id={id} [delete]
func (h *Handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusOK,
	}

	defer func() {
		w.Write(answer.getJson())
	}()

	id := r.URL.Query().Get("id")
	if id == "" {
		answer.Status = http.StatusBadRequest
		return
	}

	err := h.bookingStorage.Delete(r.Context(), id)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return
	}
}
