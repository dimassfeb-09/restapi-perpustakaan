package domain

import "time"

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Level    string    `json:"level"`
	CreateAt time.Time `json:"create_at"`
}
