package categories

type CategoriesUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
