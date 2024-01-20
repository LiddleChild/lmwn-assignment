package apperror

import "net/http"

type AppError struct {
	Name string
	Code int
}

var (
	ServiceUnavailable = &AppError{"covid-stat-server-unavailable", http.StatusServiceUnavailable}
	ResponseError      = &AppError{"covid-stat-response-error", http.StatusServiceUnavailable}
	DecodeError        = &AppError{"decode-response-error", http.StatusInternalServerError}
)
