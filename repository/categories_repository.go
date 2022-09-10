package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
)

type CategoriesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Categories) domain.Categories
	Update(ctx context.Context, tx *sql.Tx, category domain.Categories) domain.Categories
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Categories
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Categories, error)
}
