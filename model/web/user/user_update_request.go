package user

type UserUpdateRequest struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Level    string  `json:"level"`
	CreateAt []uint8 `json:"create_at"`
}
