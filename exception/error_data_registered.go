package exception

type ErrorDataRegistered struct {
	Error string `json:"error"`
}

func NewErrorDataRegistered(error string) ErrorDataRegistered {
	return ErrorDataRegistered{Error: error}
}
