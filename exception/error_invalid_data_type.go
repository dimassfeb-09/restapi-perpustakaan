package exception

type ErrorInvalidDataType struct {
	Error string `json:"error"`
}

func NewErrorInvalidDataType(err string) NotFoundError {
	return NotFoundError{
		Error: err,
	}
}
