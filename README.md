# API Documentation

## Users API
| Method 	| Documentation     			| Description		| 
| :-------- 	| :------- 				| :--------------------	|
| `GET`		| [Here](/docs/user/UserGet.md) 	| **Get All**		|
| `GET` 	| [Here](/docs/user/UserGetById.md) 	| **Get By ID** 	|
| `POST` 	| [Here](/docs/user/UserPost.md) 	| **Add Data** 		|
| `PUT` 	| [Here](/docs/user/UserPut.md) 	| **Update Data** 	|
| `DELETE` 	| [Here](/docs/user/UserDelete.md) 	| **Delete Data** 	|


## Category API
| METHOD 	| Documentation     				| Description		|
| :-------- 	| :------- 					| :--------------------	|
| `GET`		| [Here](/docs/category/CategoryGetAll.md) 	| **Get All**		|
| `GET` 	| [Here](/docs/category/CategoryGetById.md) 	| **Get By ID** 	|
| `POST` 	| [Here](/docs/category/CategoryPost.md) 	| **Add Data** 		|
| `PUT` 	| [Here](/docs/category/CategoryPut.md) 	| **Update Data** 	|
| `DELETE` 	| [Here](/docs/category/CategoryDelete.md) 	| **Delete Data** 	|

## Officer API
| METHOD 	| Documentation     				| Description		|
| :-------- 	| :------- 					| :--------------------	|
| `GET`		| [Here](/docs/officer/OfficerGetAll.md) 	| **Get All**		|
| `GET` 	| [Here](/docs/officer/OfficerGetById.md) 	| **Get By ID** 	|
| `POST` 	| [Here](/docs/officer/OfficerPost.md) 		| **Add Data** 		|
| `PUT` 	| [Here](/docs/officer/OfficerPut.md) 		| **Update Data** 	|
| `DELETE` 	| [Here](/docs/officer/OfficerDelete.md) 	| **Delete Data** 	|


## Book API
| METHOD 	| Documentation     				| Description		|
| :-------- 	| :------- 					| :--------------------	|
| `GET`		| [Here](/docs/book/BookGetAll.md) 		| **Get All**		|
| `GET` 	| [Here](/docs/book/BookGetById.md) 		| **Get By ID** 	|
| `POST` 	| [Here](/docs/book/BookPost.md) 		| **Add Data** 		|
| `PUT` 	| [Here](/docs/book/BookPut.md) 		| **Update Data** 	|
| `DELETE` 	| [Here](/docs/book/BookDelete.md) 		| **Delete Data** 	|



# Required
## Go Controller, Service, Repository
- **Controller** digunakan untuk memanajemen Interface request input oleh user untuk melanjutkan ke tahap Bisnis Logic. Lokasi: [/controller](/controller)
- **Service** digunakan untuk melakukan Middleware, Errorhandling, Validasi Data., dan Bisnis Logic untuk meneruskan ke tahap Repository. Lokasi: [/service](/service)
- **Repository** digunakan untuk memanipulasi data dengan models berinteraksi dengan Database. Lokasi: [/repository](/repository)

## Go Error Handling dan Middleware
- **Error Handling** digunakan untuk menghandle terjadinya error-error yang kemungkinan terjadi, contoh:
```go
func notFoundErr(c *gin.Context, recovered interface{}) bool {
	err, ok := recovered.(NotFoundError)

	if ok {
		webResponse := helper.webResponse(http.StatusNotFound, "Bad Not Found", err)
		c.JSON(http.StatusNotFound, webResponse)
		return true
	}
	return false
}
```
Lokasi: [/exception](/exception)

## Go Model
- **Model** digunakan untuk menyimpan data dalam bentuk objek, contoh:
```go
type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Level    string    `json:"level"`
	CreateAt time.Time `json:"create_at"`
}
```
Lokasi [/model/domain/](/model/domain/)




## Go Request, Validasi, Response
- **Request** digunakan controller untuk menyimpan data yang diinput oleh user ke dalam model, contoh:
```go
type UserCreateRequest struct {
	Name     string  `json:"name" binding:"required"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Level    string  `json:"level" binding:"required"`
}
```
Lokasi: [/model/web/user](/model/web/user)
- **Validasi** digunakan untuk memvalidasi input user sesuai dengan ketentuan
- **Response** digunakan untuk merespons atau menampilakan pesan setelah proses eksekusi berhasil, contoh:
```go
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
```
