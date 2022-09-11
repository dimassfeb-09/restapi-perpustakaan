package service

import (
	"context"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/book"
)

type BookService interface {
	Create(ctx context.Context, request book.BookCreateRequest) book.BookResponse
	Update(ctx context.Context, request book.BookUpdateRequest) book.BookResponse
	Delete(ctx context.Context, bookId int)
	FindAll(ctx context.Context) []book.BookResponse
	FindById(ctx context.Context, bookId int) book.BookResponse
}
