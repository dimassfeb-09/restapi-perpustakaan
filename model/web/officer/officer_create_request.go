package officer

type OfficerCreateRequest struct {
	Id       int    `form:"id"`
	Name     string `form:"name" binding:"required"`
	Position string `form:"position" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Address  string `form:"address" binding:"required"`
}
