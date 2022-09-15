package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"strconv"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(name, username, password, email, level) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Username, user.Password, user.Email, user.Level)
	helper.PanicIfError(err)

	lastId, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(lastId)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET id = ?, name = ?, username = ?, password = ?, email = ?, level = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Username, user.Password, user.Email, user.Level, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id, name, username, password, email, level, create_at FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Email, &user.Level, &user.CreateAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User dengan ID " + strconv.Itoa(userId) + " tidak ditemukan")
	}
}

func (repository *UserRepositoryImpl) FindByEmailCreate(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT email FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Tidak ada data dengan Email " + user.Email)
	}
}

func (repository *UserRepositoryImpl) FindByEmailUpdate(ctx context.Context, tx *sql.Tx, email string, userId int) (domain.User, error) {
	SQL := "SELECT email FROM users WHERE email = ? AND id != ?"
	rows, err := tx.QueryContext(ctx, SQL, email, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Tidak ada data dengan Email " + user.Email)
	}
}

func (repository *UserRepositoryImpl) LoginAuth(ctx context.Context, tx *sql.Tx, username string, password string) (domain.User, error) {
	SQL := "SELECT id, username FROM users WHERE username = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, SQL, username, password)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Password anda salah")
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "SELECT id, username FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Username " + username + " tidak ditemukan")
	}
}

func (repository *UserRepositoryImpl) FindByUsernameUpdate(ctx context.Context, tx *sql.Tx, username string, userId int) (domain.User, error) {
	SQL := "SELECT id, username FROM users WHERE username = ? AND id !=  ?"
	rows, err := tx.QueryContext(ctx, SQL, username, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Username " + username + " tidak ditemukan")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, name, username, email, level, create_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Level, &user.CreateAt)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}
