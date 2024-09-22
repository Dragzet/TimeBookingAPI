package handlers

import (
	"TimeBookingAPI/internal/service"
	"fmt"
	log "github.com/go-ozzo/ozzo-log"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Router  *mux.Router
	service *service.Service
	Logger  *log.Logger
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
	h.Router.ServeHTTP(writer, request)
}

func (h *Handler) RegisterUserHandlers() {
	h.Router.HandleFunc("/user", h.CreateUser).Methods(http.MethodPost)
	h.Router.HandleFunc("/user", h.DeleteUser).Methods(http.MethodDelete)
}

func (h *Handler) RegisterBookingHandlers() {
	h.Router.HandleFunc("/booking", h.CreateBooking).Methods(http.MethodPost)
	h.Router.HandleFunc("/booking", h.FindBooking).Methods(http.MethodGet)
	h.Router.HandleFunc("/booking", h.DeleteBooking).Methods(http.MethodDelete)
}

func NewHandler(router *mux.Router, services *service.Service, logger *log.Logger) *Handler {
	handler := Handler{
		Router:  router,
		service: services,
		Logger:  logger,
	}
	handler.RegisterUserHandlers()
	handler.RegisterBookingHandlers()
	handler.Router.NotFoundHandler = http.HandlerFunc(handler.DefaultHandler)
	return &handler
}

func (h *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	answer := service.Answer{
		Status: http.StatusNotFound,
	}
	w.Write(answer.GetJson())
}

func (h *Handler) LogError(r *http.Request, err error) {
	h.Logger.Error(fmt.Sprintf("Request from %s ended with error: %s", r.RemoteAddr, err.Error()))
}
