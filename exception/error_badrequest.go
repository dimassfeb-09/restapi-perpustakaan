package exception

type ErrorBadRequest struct {
	Error string `json:"error"`
}

func NewErrorBadRequest(error string) ErrorBadRequest {
	return ErrorBadRequest{Error: error}
}
