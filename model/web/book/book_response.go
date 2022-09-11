package book

type BookResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	CategoryId     int    `json:"category_id"`
	Stock          int    `json:"stock"`
	ProductsStatus string `json:"products_status"`
}
