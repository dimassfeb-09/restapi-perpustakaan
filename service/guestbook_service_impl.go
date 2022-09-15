package service

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/guestbook"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
	"time"
)

type GuestBookServiceImpl struct {
	DB                  *sql.DB
	GuestBookRepository repository.GuestBookRepository
	UserRepository      repository.UserRepository
	OfficerRepository   repository.OfficerRepository
	BookRepository      repository.BookRepository
}

func NewGuestBookServiceImpl(DB *sql.DB, guestBookRepository repository.GuestBookRepository, userRepository repository.UserRepository, officerRepository repository.OfficerRepository, bookRepository repository.BookRepository) GuestBookService {
	return &GuestBookServiceImpl{DB: DB, GuestBookRepository: guestBookRepository, UserRepository: userRepository, OfficerRepository: officerRepository, BookRepository: bookRepository}
}

func (service *GuestBookServiceImpl) Create(ctx context.Context, request guestbook.GuestBookCreateRequest) guestbook.GuestBookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	// Find User ID
	_, err = service.UserRepository.FindById(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	// Find Officer ID
	_, err = service.OfficerRepository.FindById(ctx, tx, request.OfficerId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	// Find Book ID
	_, err = service.BookRepository.FindById(ctx, tx, request.BookId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	parse, err := time.Parse("2006-01-02 15:04:05", request.EndDate)
	if err != nil {
		panic(exception.NewErrorBadRequest("Format Date Time tidak sesuai, gunakan yyyy-dd-mm HH:mm:ss"))
	} else if parse.Unix() < time.Now().Unix() {
		panic(exception.NewErrorBadRequest("End Date Time tidak boleh kurang dari hari ini"))
	}

	formatEndDate := parse.Format(time.RFC3339)
	timeParseDateEnd, _ := time.Parse(time.RFC3339, formatEndDate)

	guestBook := domain.GuestBook{
		UserId:    request.UserId,
		OfficerId: request.OfficerId,
		BookId:    request.BookId,
		EndDate:   timeParseDateEnd,
	}

	guestBookResponse := service.GuestBookRepository.Create(ctx, tx, guestBook)
	formatStartDate := time.Now().Format(time.RFC3339)
	timeParseDateStart, _ := time.Parse(time.RFC3339, formatStartDate)
	guestBookResponse.StartDate = timeParseDateStart

	return helper.ToGuestBookResponse(guestBookResponse)
}

func (service *GuestBookServiceImpl) Update(ctx context.Context, request guestbook.GuestBookUpdateRequest) guestbook.GuestBookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	// Find GuestBook ID
	findGuestBookById, err := service.GuestBookRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	// Find User ID
	_, err = service.UserRepository.FindById(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	// Find Officer ID
	_, err = service.OfficerRepository.FindById(ctx, tx, request.OfficerId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	// Find Book ID
	_, err = service.BookRepository.FindById(ctx, tx, request.BookId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	parse, err := time.Parse("2006-01-02 15:04:05", request.EndDate)
	if err != nil {
		panic(exception.NewErrorBadRequest("Format Date Time tidak sesuai, gunakan yyyy-dd-mm HH:mm:ss"))
	} else if parse.Unix() < time.Now().Unix() {
		panic(exception.NewErrorBadRequest("End Date Time tidak boleh kurang dari hari ini"))
	}
	formatEndDate := parse.Format(time.RFC3339)
	timeParseDateTime, _ := time.Parse(time.RFC3339, formatEndDate)

	guestBook := domain.GuestBook{
		Id:        request.Id,
		UserId:    request.UserId,
		OfficerId: request.OfficerId,
		BookId:    request.BookId,
		StartDate: findGuestBookById.StartDate,
		EndDate:   timeParseDateTime,
	}

	guestBookResponse := service.GuestBookRepository.Update(ctx, tx, guestBook)
	return helper.ToGuestBookResponse(guestBookResponse)
}

func (service *GuestBookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	// Find GuestBook ID
	_, err = service.GuestBookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}
}

func (service *GuestBookServiceImpl) FindAll(ctx context.Context) []guestbook.GuestBookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	// Find GuestBook All
	guestBooks := service.GuestBookRepository.FindAll(ctx, tx)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	return helper.ToGuestBookResponses(guestBooks)
}

func (service *GuestBookServiceImpl) FindById(ctx context.Context, bookId int) guestbook.GuestBookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	// Find GuestBook ID
	guestBook, err := service.GuestBookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewErrorBadRequest(err.Error()))
	}

	return helper.ToGuestBookResponse(guestBook)
}

func (service *GuestBookServiceImpl) FindByUserId(ctx context.Context, userId int) []guestbook.GuestBookResponsebyUserId {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	_, errMsg := service.UserRepository.FindById(ctx, tx, userId)
	if errMsg != nil {
		panic(exception.NewErrorNotFound(errMsg.Error()))
	}

	guestBooksByUserId := service.GuestBookRepository.FindByUserId(ctx, tx, userId)

	var guestBookResponsebyUserId []guestbook.GuestBookResponsebyUserId
	for _, data := range guestBooksByUserId {
		guestBook := guestbook.GuestBookResponsebyUserId{
			Id:           data.Id,
			Name:         data.Name,
			BookName:     data.BookName,
			CategoryName: data.CategoryName,
		}
		guestBookResponsebyUserId = append(guestBookResponsebyUserId, guestBook)
	}

	return guestBookResponsebyUserId
}
