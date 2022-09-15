package guestbook

type GuestBookResponsebyUserId struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	BookName     string `json:"book_name"`
	CategoryName string `json:"category_name"`
}
