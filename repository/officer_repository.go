package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
)

type OfficerRepository interface {
	Create(ctx context.Context, tx *sql.Tx, officer domain.Officer) domain.Officer
	Update(ctx context.Context, tx *sql.Tx, officer domain.Officer) domain.Officer
	Delete(ctx context.Context, tx *sql.Tx, officerId int)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Officer
	FindById(ctx context.Context, tx *sql.Tx, officerId int) (domain.Officer, error)
}
