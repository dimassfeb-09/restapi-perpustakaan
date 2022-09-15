package helper

import (
	"github.com/dimassfeb-09/restapi-perpustakaan/model/domain"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/book"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/categories"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/guestbook"
	"github.com/dimassfeb-09/restapi-perpustakaan/model/web/officer"
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
		ImgUrl:         domainBook.ImgUrl,
	}
}

func ToBookResponses(domainBooks []domain.Book) []book.BookResponse {
	var bookResponses []book.BookResponse
	for _, data := range domainBooks {
		bookResponses = append(bookResponses, ToBookResponse(data))
	}
	return bookResponses
}

func ToOfficerResponse(domainOfficer domain.Officer) officer.OfficerResponse {
	return officer.OfficerResponse{
		Id:       domainOfficer.Id,
		Name:     domainOfficer.Name,
		Position: domainOfficer.Position,
		Phone:    domainOfficer.Phone,
		Address:  domainOfficer.Address,
	}
}

func ToOfficerResponses(domainOfficer []domain.Officer) []officer.OfficerResponse {
	var officerResponses []officer.OfficerResponse
	for _, data := range domainOfficer {
		officerResponses = append(officerResponses, ToOfficerResponse(data))
	}
	return officerResponses
}

func ToGuestBookResponse(domainGuestBook domain.GuestBook) guestbook.GuestBookResponse {
	return guestbook.GuestBookResponse{
		Id:        domainGuestBook.Id,
		UserId:    domainGuestBook.UserId,
		BookId:    domainGuestBook.BookId,
		OfficerId: domainGuestBook.OfficerId,
		StartDate: domainGuestBook.StartDate,
		EndDate:   domainGuestBook.EndDate,
	}
}

func ToGuestBookResponses(domainGuestBook []domain.GuestBook) []guestbook.GuestBookResponse {
	var guestBookResponses []guestbook.GuestBookResponse
	for _, data := range domainGuestBook {
		guestBookResponses = append(guestBookResponses, ToGuestBookResponse(data))
	}
	return guestBookResponses
}

func ToLoginUser(domainUser domain.User) user.UserLoginResponse {
	return user.UserLoginResponse{
		Id:       domainUser.Id,
		Username: domainUser.Username,
	}
}
