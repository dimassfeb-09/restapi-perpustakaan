package service

import (
	"context"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/officer"
)

type OfficerService interface {
	Create(ctx context.Context, request officer.OfficerCreateRequest) officer.OfficerResponse
	Update(ctx context.Context, request officer.OfficerUpdateRequest) officer.OfficerResponse
	Delete(ctx context.Context, officerId int)
	FindAll(ctx context.Context) []officer.OfficerResponse
	FindById(ctx context.Context, officerId int) officer.OfficerResponse
}
