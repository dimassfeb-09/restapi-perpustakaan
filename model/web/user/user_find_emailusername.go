package user

type UserFindEmailUsername struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}
