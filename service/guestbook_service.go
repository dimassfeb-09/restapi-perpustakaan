package service

import (
	"context"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/guestbook"
)

type GuestBookService interface {
	Create(ctx context.Context, request guestbook.GuestBookCreateRequest) guestbook.GuestBookResponse
	Update(ctx context.Context, request guestbook.GuestBookUpdateRequest) guestbook.GuestBookResponse
	Delete(ctx context.Context, bookId int)
	FindAll(ctx context.Context) []guestbook.GuestBookResponse
	FindById(ctx context.Context, bookId int) guestbook.GuestBookResponse
}
