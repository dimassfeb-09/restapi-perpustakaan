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

type BookRepositoryImpl struct {
}

func NewBookRepositoryImpl() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "INSERT INTO books(name, category_id, stock, products_status, img_url) VALUES (?,?,?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, book.Name, book.CategoryId, book.Stock, book.ProductsStatus, book.ImgUrl)
	helper.PanicIfError(err)

	lastInsertId, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(lastInsertId)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "UPDATE books SET name = ?, category_id = ?, stock = ?, products_status = ?, img_url = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Name, book.CategoryId, book.Stock, book.ProductsStatus, book.ImgUrl, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) UpdateStock(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "UPDATE books SET stock = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Stock, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bookId int) {
	SQL := "DELETE FROM books WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, bookId)
	if err != nil {
		panic(exception.NewErrorForbidden("Tidak dapat menghapus Buku yang dipinjam"))
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "SELECT id, name, category_id, create_at, stock, products_status, img_url FROM books"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.Id, &book.Name, &book.CategoryId, &book.CreateAt, &book.Stock, &book.ProductsStatus, &book.ImgUrl)
		helper.PanicIfError(err)
		books = append(books, book)
	}

	return books
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	SQL := "SELECT id, name, category_id, create_at, stock, products_status, img_url FROM books WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	var book domain.Book
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Name, &book.CategoryId, &book.CreateAt, &book.Stock, &book.ProductsStatus, &book.ImgUrl)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("Buku dengan ID " + strconv.Itoa(bookId) + " tidak ditemukan")
	}
}
