package service

import (
	"context"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/categories"
)

type CategoriesService interface {
	Create(ctx context.Context, request categories.CategoriesCreateRequest) categories.CategoriesResponse
	Update(ctx context.Context, request categories.CategoriesUpdateRequest) categories.CategoriesResponse
	Delete(ctx context.Context, categoryId int)
	FindAll(ctx context.Context) []categories.CategoriesResponse
	FindById(ctx context.Context, categoryId int) categories.CategoriesResponse
}
