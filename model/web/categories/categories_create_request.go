package categories

type CategoriesCreateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
