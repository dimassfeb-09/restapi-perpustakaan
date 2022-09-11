package book

type BookUpdateRequest struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	CategoryId     int    `json:"category_id"`
	Stock          int    `json:"stock"`
	ProductsStatus string `json:"products_status"`
}
