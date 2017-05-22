package errors

type (
	ApiErrors struct {
		Errors []ApiError `json:"errors"`
	}

	ApiError struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}
)

func NewApiErrors() *ApiErrors {
	return &ApiErrors{}
}

func (e *ApiErrors) AddError(status int, msg string) *ApiErrors {
	e.Errors = append(e.Errors, ApiError{
		StatusCode: status,
		Message:    msg,
	})
	return e
}
