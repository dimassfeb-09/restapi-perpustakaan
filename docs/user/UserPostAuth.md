## POST User Auth
- End Point URL: `/user/login`
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
      | `username` | `string` | **Required**. Your Username account |
      | `password` | `string` | **Required**. Your Password account |
    - Response Body
      ```json
      {
          "code": 200,
          "status": "OK",
          "data": {
              "id": "int",
              "username": "string",
          }
      }
      ```
    - Response Error
        - Bad Not Found
      ```json
      {
          "code": 404,
          "status": "Bad Not Found",
          "data": {
              "error": "Username tidak ditemukan",
          }
      }
      ```
      ```json
      {
          "code": 404,
          "status": "Bad Not Found",
          "data": {
              "error": "Password anda salah"
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
