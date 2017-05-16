package errors

type (
	APIErrors struct {
		Errors []ApiError `json:"errors"`
	}

	ApiError struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}
)
