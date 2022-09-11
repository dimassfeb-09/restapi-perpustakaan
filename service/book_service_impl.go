package service

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/book"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
)

type BookServiceImpl struct {
	DB                 *sql.DB
	BookRepository     repository.BookRepository
	CategoryRepository repository.CategoriesRepository
}

func NewBookServiceImpl(DB *sql.DB, bookRepository repository.BookRepository, categoryRepository repository.CategoriesRepository) BookService {
	return &BookServiceImpl{DB: DB, BookRepository: bookRepository, CategoryRepository: categoryRepository}
}

func (service *BookServiceImpl) Create(ctx context.Context, request book.BookCreateRequest) book.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	if request.ProductsStatus != "in_stock" && request.ProductsStatus != "out_of_stock" {
		panic(exception.NewErrorBadRequest("Produk status harus in_stock atau out_of_stock"))
	}

	_, errMsg := service.CategoryRepository.FindById(ctx, tx, request.CategoryId)
	if errMsg != nil {
		panic(exception.NewErrorBadRequest(errMsg.Error()))
	}

	domainBook := domain.Book{
		Name:           request.Name,
		CategoryId:     request.CategoryId,
		Stock:          request.Stock,
		ProductsStatus: request.ProductsStatus,
	}

	createResponse := service.BookRepository.Create(ctx, tx, domainBook)

	return helper.ToBookResponse(createResponse)
}

func (service *BookServiceImpl) Update(ctx context.Context, request book.BookUpdateRequest) book.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, errMsg := service.CategoryRepository.FindById(ctx, tx, request.CategoryId)
	if errMsg != nil {
		panic(exception.NewErrorBadRequest(errMsg.Error()))
	}

	domainBook := domain.Book{
		Id:             request.Id,
		Name:           request.Name,
		CategoryId:     request.CategoryId,
		Stock:          request.Stock,
		ProductsStatus: request.ProductsStatus,
	}

	createResponse := service.BookRepository.Update(ctx, tx, domainBook)

	return helper.ToBookResponse(createResponse)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, bookId)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []book.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	bookResponse := service.BookRepository.FindAll(ctx, tx)

	return helper.ToBookResponses(bookResponse)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) book.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	bookResponse, err := service.BookRepository.FindById(ctx, tx, bookId)
	helper.PanicIfError(err)

	return helper.ToBookResponse(bookResponse)
}
