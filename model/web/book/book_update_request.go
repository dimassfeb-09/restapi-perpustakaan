package book

type BookUpdateRequest struct {
	Id             int    `form:"id"`
	Name           string `form:"name"`
	CategoryId     int    `form:"category_id"`
	Stock          int    `form:"stock"`
	ProductsStatus string `form:"products_status"`
}
