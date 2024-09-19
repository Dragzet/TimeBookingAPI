package handlers

import (
	"TimeBookingAPI/internal/userModule"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {

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

	user, err := h.userStorage.Find(r.Context(), username)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return
	}

	answer.Data = user
}

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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	updatedUser := userModule.New()
	err = json.Unmarshal(d, &updatedUser)
	if err != nil {
		answer.Status = http.StatusBadRequest
		return
	}

	fmt.Println(updatedUser)
	err = h.userStorage.Update(r.Context(), updatedUser)
	if err != nil {
		fmt.Println(err)
		answer.Status = http.StatusInternalServerError
		return
	}
}
