## POST User
- End Point URL: `/user/add`
    - Method: `POST`
    - Content-Type: `multipart/form-data`
    - Accept: `multipart/form-data`
    - Header:
      | Header 	| Type     | Description                |
      | :-------- | :------- | :------------------------- |
      | `X-API-KEY` | `string` | **Required**. Your API key |
    - Post Form:
        | Key 	| Value     | Description                |
        | :-------- | :------- | :------------------------- |
        | `name` | `string` | **Required**. Full Name |
        | `username` | `string` | **Required**. Username |
        | `password` | `string` | **Required**. Password |
        | `email` | `string` | **Required**. Email |
        | `level` | `string` | **Required**. Level |
    - Response Error
        - Bad Not Found
      ```json
      {
          "code": 404,
          "status": "Bad Not Found",
          "data": {
              "error": "Data tidak ditemukan",
          }
      }
      ```
        - Bad Request
      ```json
      {
          "code": 400,
          "status": "Bad Request",
          "data": {
              "error": "Username telah terdaftar"
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