package handlers

import (
	"TimeBookingAPI/internal/bookingModule"
	"TimeBookingAPI/internal/userModule"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Answer представляет собой ответ API с кодом состояния и данными.
type Answer struct {
	Status int         `json:"status"` // Код состояния HTTP
	Data   interface{} `json:"data"`   // Данные ответа
}

// getJson сериализует ответ в формат JSON.
func (answer *Answer) getJson() []byte {
	d, err := json.Marshal(answer)
	if err != nil {
		return []byte{}
	}
	return d
}

// Handler обрабатывает HTTP-запросы и маршрутизацию для API.
type Handler struct {
	router         *mux.Router                  // Маршрутизатор для обработки HTTP-запросов
	bookingStorage bookingModule.BookingStorage // Хранилище для бронирований
	userStorage    userModule.UserStorage       // Хранилище для пользователей
	UserHandler                                 // Интерфейс для работы с пользователями
	BookingHandler                              // Интерфейс для работы с бронированиями
}

// UserHandler определяет методы для работы с пользователями.
type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) // Создает нового пользователя
	DeleteUser(w http.ResponseWriter, r *http.Request) // Удаляет пользователя
}

// BookingHandler определяет методы для работы с бронированиями.
type BookingHandler interface {
	FindBooking(w http.ResponseWriter, r *http.Request)   // Находит бронирование по ID
	CreateBooking(w http.ResponseWriter, r *http.Request) // Создает новое бронирование
	DeleteBooking(w http.ResponseWriter, r *http.Request) // Удаляет бронирование
}

// ServeHTTP обрабатывает HTTP-запросы.
func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.router.ServeHTTP(writer, request)
}

// RegisterUserHandlers регистрирует маршруты для обработки запросов пользователей.
func (h *Handler) RegisterUserHandlers() {
	h.router.HandleFunc("/user", h.CreateUser).Methods(http.MethodPost)   // Создание пользователя
	h.router.HandleFunc("/user", h.DeleteUser).Methods(http.MethodDelete) // Удаление пользователя
}

// RegisterBookingHandlers регистрирует маршруты для обработки запросов бронирования.
func (h *Handler) RegisterBookingHandlers() {
	h.router.HandleFunc("/booking", h.CreateBooking).Methods(http.MethodPost)   // Создание бронирования
	h.router.HandleFunc("/booking", h.FindBooking).Methods(http.MethodGet)      // Поиск бронирования
	h.router.HandleFunc("/booking", h.DeleteBooking).Methods(http.MethodDelete) // Удаление бронирования
}

// NewHandler создает новый обработчик с зарегистрированными маршрутами.
func NewHandler(router *mux.Router, bookingStorage bookingModule.BookingStorage, userStorage userModule.UserStorage) *Handler {
	handler := Handler{
		router:         router,
		bookingStorage: bookingStorage,
		userStorage:    userStorage,
	}
	handler.RegisterUserHandlers()
	handler.RegisterBookingHandlers()
	handler.router.NotFoundHandler = http.HandlerFunc(handler.DefaultHandler) // Обработчик для 404
	return &handler
}

// DefaultHandler обрабатывает запросы к несуществующим маршрутам.
func (h *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	answer := Answer{
		Status: http.StatusNotFound,
	}

	w.Write(answer.getJson())
}
