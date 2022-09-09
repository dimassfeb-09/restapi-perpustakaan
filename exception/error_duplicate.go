package exception

type ErrorDuplicate struct {
	Error string `json:"error"`
}

func NewErrorDuplicate(error string) ErrorDuplicate {
	return ErrorDuplicate{Error: error}
}
