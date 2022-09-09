package user

import "time"

type UserResponse struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Level    string    `json:"level"`
	CreateAt time.Time `json:"create_at"`
}
