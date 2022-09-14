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
