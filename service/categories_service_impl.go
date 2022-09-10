package service

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/categories"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
)

type CategoryServiceImpl struct {
	CategoriesRepository repository.CategoriesRepository
	DB                   *sql.DB
}

func NewCategoryServiceImpl(categoriesRepository repository.CategoriesRepository, DB *sql.DB) CategoriesService {
	return &CategoryServiceImpl{CategoriesRepository: categoriesRepository, DB: DB}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request categories.CategoriesCreateRequest) categories.CategoriesResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := domain.Categories{
		Name: request.Name,
	}

	categoryResponse := service.CategoriesRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(categoryResponse)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request categories.CategoriesUpdateRequest) categories.CategoriesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.CategoriesRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	category := domain.Categories{
		Id:   request.Id,
		Name: request.Name,
	}

	categoryResponse := service.CategoriesRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(categoryResponse)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.CategoriesRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.CategoriesRepository.Delete(ctx, tx, categoryId)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []categories.CategoriesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	categoriesResponse := service.CategoriesRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categoriesResponse)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) categories.CategoriesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	_, err = service.CategoriesRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	categoryResponse, err := service.CategoriesRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(categoryResponse)
}
