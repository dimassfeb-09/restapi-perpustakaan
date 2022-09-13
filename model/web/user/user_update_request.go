package user

import "time"

type UserUpdateRequest struct {
	Id       int       `form:"id"`
	Name     string    `form:"name" binding:"required"`
	Username string    `form:"username" binding:"required"`
	Password string    `form:"password" binding:"required"`
	Email    string    `form:"email" binding:"required"`
	Level    string    `form:"level" binding:"required"`
	CreateAt time.Time `form:"create_at"`
}
