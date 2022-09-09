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
	FindBy(ctx context.Context, tx *sql.Tx, filterBy string, value interface{}) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
