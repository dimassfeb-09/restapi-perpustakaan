package exception

type ErrorShouldBind struct {
	Error string `json:"error"`
}

func NewErrorShouldBind(err string) ErrorShouldBind {
	return ErrorShouldBind{
		Error: err,
	}
}
