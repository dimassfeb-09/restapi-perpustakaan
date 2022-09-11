package exception

type ErrorForbidden struct {
	Error string `json:"error"`
}

func NewErrorForbidden(error string) ErrorForbidden {
	return ErrorForbidden{Error: error}
}
