package apperror

type BizError struct {
	Code    int
	Message string
}

func (e *BizError) Error() string {
	return e.Message
}

func New(code int, msg string) *BizError {
	return &BizError{
		Code:    code,
		Message: msg,
	}
}
