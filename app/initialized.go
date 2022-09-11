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
