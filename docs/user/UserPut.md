## PUT User
- End Point URL: `/user/update/{userId}`
    - Method: `PUT`
    - Content-Type: `application/json`
    - Accept: `application/json`
    - Header:
      | Header 	| Type     | Description                |
      | :-------- | :------- | :------------------------- |
      | `X-API-KEY` | `string` | **Required**. Your API key |
    - Body:
      ```json
      {
          "name": "string",
          "username": "string",
          "password": "string",
          "email": "string",
          "level": "string"
      }
      ```
    - Response Body
      ```json
      {
          "code": 200,
          "status": "OK",
          "data": {
              "id": "int",
              "name": "string",
              "username": "string",
              "email": "string",
              "level": "string",
              "create_at": "datetime"
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
