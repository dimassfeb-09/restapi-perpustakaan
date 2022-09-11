package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"strconv"
)

type CategoriesRepositoryImpl struct {
}

func NewCategoriesRepositoryImpl() CategoriesRepository {
	return &CategoriesRepositoryImpl{}
}

func (repository *CategoriesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Categories) domain.Categories {
	SQL := "INSERT INTO categories(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	lastInsertId, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(lastInsertId)
	return category
}

func (repository *CategoriesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Categories) domain.Categories {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoriesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	if err != nil {
		panic(exception.NewErrorForbidden("Tidak dapat menghapus kategori yang terdapat buku"))
	}
}

func (repository *CategoriesRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Categories {
	SQL := "SELECT id, name FROM categories ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Categories
	for rows.Next() {
		var category domain.Categories
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoriesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Categories, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var category domain.Categories
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		Id := strconv.Itoa(categoryId)
		return category, errors.New("Category dengan ID " + Id + " tidak ditemukan")
	}

}
