package handlers

import (
	"TimeBookingAPI/internal/userModule"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user with the provided information
// @Tags user
// @Accept json
// @Produce json
// @Param user body userModule.UserModel true "User data"
// @Success 201 {object} Answer
// @Failure 400 {object} Answer "Invalid input"
// @Failure 500 {object} Answer "Internal server error"
// @Router /user [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

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

	newUser := userModule.New()
	err = json.Unmarshal(d, &newUser)
	if err != nil {
		answer.Status = http.StatusBadRequest
		return
	}

	if len(newUser.Username) < 5 || len(newUser.Password) < 5 || len(newUser.Password) > 72 {
		answer.Status = http.StatusBadRequest
		return
	}

	err = h.userStorage.Create(r.Context(), newUser)
	if err != nil {
		fmt.Println(err.Error())
		answer.Status = http.StatusInternalServerError
		return
	}
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {object} Answer "No content"
// @Failure 404 {object} Answer "User not found"
// @Failure 500 {object} Answer "Internal server error"
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	err := h.userStorage.Delete(r.Context(), h.bookingStorage, username)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return
	}
}
