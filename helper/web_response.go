package helper

import "github.com/dimassfeb-09/restapi-perpustakaan/model/web"

func WebResponse(code int, status string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
