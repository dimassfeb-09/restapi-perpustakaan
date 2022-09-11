package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
)

type BookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Delete(ctx context.Context, tx *sql.Tx, bookId int)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Book
	FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error)
}
