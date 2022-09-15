package categories

type CategoriesUpdateRequest struct {
	Id   int    `form:"id"`
	Name string `form:"name" binding:"required"`
}
