package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
)

type GuestBookRepository interface {
	Create(ctx context.Context, tx *sql.Tx, book domain.GuestBook) domain.GuestBook
	Update(ctx context.Context, tx *sql.Tx, book domain.GuestBook) domain.GuestBook
	FindAll(ctx context.Context, tx *sql.Tx) []domain.GuestBook
	FindById(ctx context.Context, tx *sql.Tx, guestBookId int) (domain.GuestBook, error)
	Delete(ctx context.Context, tx *sql.Tx, guestBookId int)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.GuestBookByUserId
}
