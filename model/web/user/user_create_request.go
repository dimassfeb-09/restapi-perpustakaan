package user

type UserCreateRequest struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Level    string  `json:"level"`
	CreateAt []uint8 `json:"create_at"`
}
