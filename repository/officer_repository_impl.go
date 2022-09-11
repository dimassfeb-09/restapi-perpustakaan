package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"strconv"
)

type OfficerRepositoryImpl struct {
}

func NewOfficerRepositoryImpl() OfficerRepository {
	return &OfficerRepositoryImpl{}
}

func (repository *OfficerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, officer domain.Officer) domain.Officer {
	SQL := "INSERT INTO officers(name, position, phone, address) VALUES (?,?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, officer.Name, officer.Position, officer.Phone, officer.Address)
	helper.PanicIfError(err)

	lastInsertId, err := result.LastInsertId()
	helper.PanicIfError(err)

	officer.Id = int(lastInsertId)
	return officer
}

func (repository *OfficerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, officer domain.Officer) domain.Officer {
	SQL := "UPDATE officers SET name = ?, position = ?, phone = ?, address = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, officer.Name, officer.Position, officer.Phone, officer.Address, officer.Id)
	helper.PanicIfError(err)

	return officer
}

func (repository *OfficerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, officerId int) {
	SQL := "DELETE FROM officers WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, officerId)
	helper.PanicIfError(err)
}

func (repository *OfficerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Officer {
	SQL := "SELECT id, name, position, phone, address FROM officers"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var officers []domain.Officer
	for rows.Next() {
		var officer domain.Officer
		err := rows.Scan(&officer.Id, &officer.Name, &officer.Position, &officer.Phone, &officer.Address)
		helper.PanicIfError(err)
		officers = append(officers, officer)
	}

	return officers
}

func (repository *OfficerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, officerId int) (domain.Officer, error) {
	SQL := "SELECT id, name, position, phone, address FROM officers WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, officerId)
	helper.PanicIfError(err)
	defer rows.Close()

	var officer domain.Officer
	if rows.Next() {
		err := rows.Scan(&officer.Id, &officer.Name, &officer.Position, &officer.Phone, &officer.Address)
		helper.PanicIfError(err)
		return officer, nil
	} else {
		return officer, errors.New("Officer dengan ID " + strconv.Itoa(officerId) + " tidak ditemukan")
	}
}
