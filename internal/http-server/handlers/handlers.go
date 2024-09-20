package handlers

import (
	"TimeBookingAPI/internal/bookingModule"
	"TimeBookingAPI/internal/userModule"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Answer struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (answer *Answer) getJson() []byte {
	d, err := json.Marshal(answer)
	if err != nil {
		return []byte{}
	}
	return d
}

type Handler struct {
	router         *mux.Router
	bookingStorage bookingModule.BookingStorage
	userStorage    userModule.UserStorage
	UserHandler
	BookingHandler
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type BookingHandler interface {
	FindBooking(w http.ResponseWriter, r *http.Request)
	CreateBooking(w http.ResponseWriter, r *http.Request)
	DeleteBooking(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.router.ServeHTTP(writer, request)
}

func (h *Handler) RegisterUserHandlers() {
	h.router.HandleFunc("/user", h.CreateUser).Methods(http.MethodPost)
	h.router.HandleFunc("/user", h.DeleteUser).Methods(http.MethodDelete)

}

func (h *Handler) RegisterBookingHandlers() {
	h.router.HandleFunc("/booking", h.CreateBooking).Methods(http.MethodPost)
	h.router.HandleFunc("/booking", h.FindBooking).Methods(http.MethodGet)
	h.router.HandleFunc("/booking", h.DeleteBooking).Methods(http.MethodDelete)
}

func NewHandler(router *mux.Router, bookingStorage bookingModule.BookingStorage, userStorage userModule.UserStorage) *Handler {
	handler := Handler{
		router:         router,
		bookingStorage: bookingStorage,
		userStorage:    userStorage,
	}
	handler.RegisterUserHandlers()
	handler.RegisterBookingHandlers()
	handler.router.NotFoundHandler = http.HandlerFunc(handler.DefaultHandler)
	return &handler
}

func (h *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusNotFound,
	}

	w.Write(answer.getJson())
}
