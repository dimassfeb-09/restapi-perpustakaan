package guestbook

import "time"

type GuestBookResponse struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id" binding:"required"`
	OfficerId int       `json:"officer_id" binding:"required"`
	BookId    int       `json:"book_id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}
