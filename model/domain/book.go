package domain

import "time"

type Book struct {
	Id             int       `json:"id"`
	Name           int       `json:"name"`
	CategoryId     int       `json:"category_id"`
	CreateAt       time.Time `json:"create_at"`
	Stock          int       `json:"stock"`
	ProductsStatus string    `json:"products_status"`
}
