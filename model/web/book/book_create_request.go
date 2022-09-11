package book

type BookCreateRequest struct {
	Id             int    `json:"id" form:"id"`
	Name           string `json:"name" binding:"required" form:"name"`
	CategoryId     int    `json:"category_id" binding:"required" form:"category_id"`
	Stock          int    `json:"stock" binding:"required" form:"stock"`
	ProductsStatus string `json:"products_status" binding:"required" form:"products_status"`
}
