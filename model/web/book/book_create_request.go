package book

type BookCreateRequest struct {
	Id             int    `form:"id"`
	Name           string `binding:"required" form:"name"`
	CategoryId     int    `binding:"required" form:"category_id"`
	Stock          int    `binding:"required" form:"stock"`
	ProductsStatus string `binding:"required" form:"products_status"`
}
