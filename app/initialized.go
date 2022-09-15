package app

import (
	"database/sql"
	"github.com/dimassfeb-09/restapi-perpustakaan/controller"
	"github.com/dimassfeb-09/restapi-perpustakaan/repository"
	"github.com/dimassfeb-09/restapi-perpustakaan/service"
)

func NewInitializedUser(db *sql.DB) controller.UserController {
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(db, userRepository)
	userController := controller.NewUserControllerImpl(userService)

	return userController
}

func NewInitializedCategories(db *sql.DB) controller.CategoriesController {
	categoriesRepository := repository.NewCategoriesRepositoryImpl()
	categoriesService := service.NewCategoryServiceImpl(categoriesRepository, db)
	categoriesController := controller.NewCategoriesControllerImpl(categoriesService)

	return categoriesController
}

func NewInitializedBook(db *sql.DB) controller.BookController {
	categoriesRepository := repository.NewCategoriesRepositoryImpl()

	bookRepository := repository.NewBookRepositoryImpl()
	bookService := service.NewBookServiceImpl(db, bookRepository, categoriesRepository)
	bookController := controller.NewBookControllerImpl(bookService)

	return bookController
}

func NewInitializedOfficer(db *sql.DB) controller.OfficerController {
	categoriesRepository := repository.NewCategoriesRepositoryImpl()

	officerRepository := repository.NewOfficerRepositoryImpl()
	officerService := service.NewOfficerServiceImpl(db, officerRepository, categoriesRepository)
	officerController := controller.NewOfficerControllerImpl(officerService)

	return officerController
}

func NewInitializedGuestBook(db *sql.DB) controller.GuestBookController {
	userRepository := repository.NewUserRepositoryImpl()
	officerRepository := repository.NewOfficerRepositoryImpl()
	bookRepository := repository.NewBookRepositoryImpl()

	guestbookRepository := repository.NewGuestBookRepositoryImpl()
	guestbookService := service.NewGuestBookServiceImpl(db, guestbookRepository, userRepository, officerRepository, bookRepository)
	guestbookController := controller.NewGuestBookControllerImpl(guestbookService)

	return guestbookController
}
