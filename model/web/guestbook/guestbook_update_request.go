package guestbook

type GuestBookUpdateRequest struct {
	Id        int    `form:"id"`
	UserId    int    `form:"user_id" binding:"required"`
	OfficerId int    `form:"officer_id" binding:"required"`
	BookId    int    `form:"book_id" binding:"required"`
	EndDate   string `form:"end_date"`
}
