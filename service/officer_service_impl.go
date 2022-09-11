package service

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/exception"
	"github.com/dimassfeb-09/restapi-perpustakaan/helper"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/officer"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
)

type OfficerServiceImpl struct {
	DB                *sql.DB
	OfficerRepository repository.OfficerRepository
}

func NewOfficerServiceImpl(DB *sql.DB, officerRepository repository.OfficerRepository, categoryRepository repository.CategoriesRepository) OfficerService {
	return &OfficerServiceImpl{DB: DB, OfficerRepository: officerRepository}
}

func (service *OfficerServiceImpl) Create(ctx context.Context, request officer.OfficerCreateRequest) officer.OfficerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	domainOfficer := domain.Officer{
		Name:     request.Name,
		Position: request.Position,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	createResponse := service.OfficerRepository.Create(ctx, tx, domainOfficer)

	return helper.ToOfficerResponse(createResponse)
}

func (service *OfficerServiceImpl) Update(ctx context.Context, request officer.OfficerUpdateRequest) officer.OfficerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, errMsg := service.OfficerRepository.FindById(ctx, tx, request.Id)
	if errMsg != nil {
		panic(exception.NewErrorBadRequest(errMsg.Error()))
	}

	domainOfficer := domain.Officer{
		Id:       request.Id,
		Name:     request.Name,
		Position: request.Position,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	createResponse := service.OfficerRepository.Update(ctx, tx, domainOfficer)

	return helper.ToOfficerResponse(createResponse)
}

func (service *OfficerServiceImpl) Delete(ctx context.Context, officerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	_, err = service.OfficerRepository.FindById(ctx, tx, officerId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	service.OfficerRepository.Delete(ctx, tx, officerId)
}

func (service *OfficerServiceImpl) FindAll(ctx context.Context) []officer.OfficerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	officerResponse := service.OfficerRepository.FindAll(ctx, tx)

	return helper.ToOfficerResponses(officerResponse)
}

func (service *OfficerServiceImpl) FindById(ctx context.Context, officerId int) officer.OfficerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	officerResponse, err := service.OfficerRepository.FindById(ctx, tx, officerId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return helper.ToOfficerResponse(officerResponse)
}
