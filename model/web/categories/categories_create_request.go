package categories

type CategoriesCreateRequest struct {
	Id   int    `form:"id"`
	Name string `form:"name" binding:"required"`
}
