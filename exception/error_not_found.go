package exception

type NotFoundError struct {
	Error string `json:"error"`
}

func NewErrorNotFound(err string) NotFoundError {
	return NotFoundError{
		Error: err,
	}
}
