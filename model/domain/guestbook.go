package domain

import "time"

type GuestBook struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	OfficerId int       `json:"officer_id"`
	BookId    int       `json:"book_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type GuestBookByUserId struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	BookName     string `json:"book_name"`
	CategoryName string `json:"category_name"`
}
