package officer

type OfficerCreateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Position string `json:"position" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
