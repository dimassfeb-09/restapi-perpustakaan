package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	LoginAuth(ctx context.Context, tx *sql.Tx, username string, password string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
