package helper

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/book"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/categories"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/user"
)

func ToUserResponse(domainUser domain.User) user.UserResponse {
	return user.UserResponse{
		Id:       domainUser.Id,
		Name:     domainUser.Name,
		Username: domainUser.Username,
		Email:    domainUser.Email,
		Level:    domainUser.Level,
	}
}

func ToUserResponses(domainUsers []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, data := range domainUsers {
		userResponses = append(userResponses, ToUserResponse(data))
	}

	return userResponses
}

func ToCategoryResponse(domainCategory domain.Categories) categories.CategoriesResponse {
	return categories.CategoriesResponse{
		Id:   domainCategory.Id,
		Name: domainCategory.Name,
	}
}

func ToCategoryResponses(domainCategories []domain.Categories) []categories.CategoriesResponse {
	var categoryResponses []categories.CategoriesResponse
	for _, data := range domainCategories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(data))
	}

	return categoryResponses
}

func ToBookResponse(domainBook domain.Book) book.BookResponse {
	return book.BookResponse{
		Id:             domainBook.Id,
		Name:           domainBook.Name,
		CategoryId:     domainBook.CategoryId,
		Stock:          domainBook.Stock,
		ProductsStatus: domainBook.ProductsStatus,
	}
}

func ToBookResponses(domainBooks []domain.Book) []book.BookResponse {
	var bookResponses []book.BookResponse
	for _, data := range domainBooks {
		bookResponses = append(bookResponses, ToBookResponse(data))
	}

	return bookResponses
}
