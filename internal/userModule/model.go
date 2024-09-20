package userModule

import (
	"time"
)

// UserModel представляет собой модель пользователя в системе.
type UserModel struct {
	ID        string    `json:"ID"`        // Уникальный идентификатор пользователя
	Username  string    `json:"username"`  // Имя пользователя (логин)
	Password  string    `json:"password"`  // Пароль пользователя (в идеале, храните в хешированном виде)
	CreatedAt time.Time `json:"createdAt"` // Время создания записи пользователя
	UpdatedAt time.Time `json:"updatedAt"` // Время последнего обновления записи пользователя
}

// New создает и возвращает новую модель пользователя с текущими временными метками.
func New() *UserModel {
	return &UserModel{
		CreatedAt: time.Now(), // Устанавливаем время создания
		UpdatedAt: time.Now(), // Устанавливаем время обновления
	}
}
