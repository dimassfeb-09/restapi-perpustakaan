## PUT Officer
- End Point URL: `/officer/update/{officerId}`
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
          "position": "string",
          "phone": "string",
          "address": "string"
      }
      ```
    - Response Body
      ```json
      {
          "code": 200,
          "status": "OK",
          "data": {
              "id": "int",
              "name": "string"
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
              "error": "Officer dengan ID 10 tidak ditemukan"
          }
      }
      ```
        - Bad Request
      ```json
      {
          "code": 400,
          "status": "Bad Request",
          "data": {
              "error": "Officer tidak boleh kosong"
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