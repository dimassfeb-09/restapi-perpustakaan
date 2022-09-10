## GET Users
- End Point URL: `/user`
    - Method: `GET`
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
         "data": [
           {
             "id": "int",
             "name": "string",
             "username": "string",
             "email": "string",
             "level": "string",
             "create_at": "datetime"
           },
           {
             "id": "int",
             "name": "string",
             "username": "string",
             "email": "string",
             "level": "string",
             "create_at": "datetime"
          }
        ]
      }
      ```
    - Response Error
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