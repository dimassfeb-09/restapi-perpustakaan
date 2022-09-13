## GET All Categories
- End Point URL: `/category`
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
              "name": "string"
            },
            {
              "id": "int",
              "name": "string"
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