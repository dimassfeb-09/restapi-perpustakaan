package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"strconv"
	"time"
)

type GuestBookRepositoryImpl struct {
}

func NewGuestBookRepositoryImpl() GuestBookRepository {
	return &GuestBookRepositoryImpl{}
}

func (repository *GuestBookRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, book domain.GuestBook) domain.GuestBook {
	SQL := "INSERT INTO guestbook(user_Id, officer_id, book_id, start_date, end_date) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, book.UserId, book.OfficerId, book.BookId, time.Now(), book.EndDate)
	helper.PanicIfError(err)

	lastInsertId, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(lastInsertId)

	return book
}

func (repository *GuestBookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.GuestBook) domain.GuestBook {
	SQL := "UPDATE guestbook SET user_Id = ?, officer_id = ?, book_id = ?, end_date = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.UserId, book.OfficerId, book.BookId, book.EndDate, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *GuestBookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.GuestBook {
	SQL := "SELECT id, user_id, officer_id, book_id, start_date, end_date FROM guestbook ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var guestBooks []domain.GuestBook
	for rows.Next() {
		book := domain.GuestBook{}
		err := rows.Scan(&book.Id, &book.UserId, &book.OfficerId, &book.BookId, &book.StartDate, &book.EndDate)
		helper.PanicIfError(err)

		guestBooks = append(guestBooks, book)
	}

	return guestBooks
}

func (repository *GuestBookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, guestBookId int) (domain.GuestBook, error) {
	SQL := "SELECT id, user_id, officer_id, book_id, start_date, end_date FROM guestbook WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, guestBookId)
	helper.PanicIfError(err)
	defer rows.Close()

	var guestBook domain.GuestBook
	if rows.Next() {
		err := rows.Scan(&guestBook.Id, &guestBook.UserId, &guestBook.OfficerId, &guestBook.BookId, &guestBook.StartDate, &guestBook.EndDate)
		helper.PanicIfError(err)
		return guestBook, nil
	} else {
		guestIdStr := strconv.Itoa(guestBookId)
		return guestBook, errors.New("Peminjam dengan ID " + guestIdStr + " tidak ditemukan")
	}
}

func (repository *GuestBookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, guestBookId int) {
	SQL := "DELETE FROM guestbook WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, guestBookId)
	helper.PanicIfError(err)
}

func (repository *GuestBookRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.GuestBookByUserId {
	SQL := `SELECT guestbook.id as id, user.name as name, book.name as book_name, category.name as category_name
    FROM guestbook
        JOIN users user on user.id = guestbook.user_id
            JOIN books book on book.id = guestbook.book_id
                JOIN categories category on book.category_id = category.id
                    WHERE user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var guestBooks []domain.GuestBookByUserId
	for rows.Next() {
		var guestBook domain.GuestBookByUserId
		err := rows.Scan(&guestBook.Id, &guestBook.Name, &guestBook.BookName, &guestBook.CategoryName)
		helper.PanicIfError(err)

		guestBooks = append(guestBooks, guestBook)
	}

	return guestBooks
}
