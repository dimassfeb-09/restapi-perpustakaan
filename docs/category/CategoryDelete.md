## DELETE Category
- End Point URL: `/category/delete/{categoryId}`
    - Method: `DELETE`
    - Content-Type: `application/json`
    - Accept: `application/json`
    - Header:
      | Header 	| Type     | Description                |
      | :-------- | :------- | :------------------------- |
      | `X-API-KEY` | `string` | **Required**. Your API key |
    - Response Body
      ```json
      {
          "code": 200,
          "status": "OK",
          "data": "Category Berhasil dihapus"
      }
      ```
    - Response Error
        - Bad Not Found
      ```json
      {
          "code": 404,
          "status": "Bad Not Found",
          "data": {
              "error": "Kategori dengan ID 10 tidak ditemukan"
          }
      }
      ```
        - Internal Server Error
      ```json
      {
          "code": 500,
          "status": "Status Internal Server Error",
          "data": {
              "error": "Status Internal Server Error"
          }
      }
      ```
        - Forbidden
      ```json
      {
          "code": 403,
          "status": "Forbidden",
          "data": {
              "error": "Kategori tidak dapat dihapus, karena terdapat produk"
          }
      }
      ```