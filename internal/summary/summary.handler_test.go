package summary

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LiddleChild/covid-stat/apperror"
	"github.com/gin-gonic/gin"
)

type mockService struct {
	getSummary func(*Summary) *apperror.AppError
}

func (h *mockService) GetSummary(summary *Summary) *apperror.AppError {
	return h.getSummary(summary)
}

func TestGetSummary(t *testing.T) {
	testCases := []struct {
		name         string
		err          *apperror.AppError
		expectedCode int
	}{
		{
			name:         "OK",
			err:          nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Internal Server Error",
			err:          apperror.ServiceUnavailable,
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := mockService{
				getSummary: func(s *Summary) *apperror.AppError {
					return tc.err
				},
			}

			handler := NewHandler(&service)

			router := gin.New()
			router.GET("/covid/summary", handler.GetSummary)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/covid/summary", nil)
			router.ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Response code is wrong.\nError: %v\nResult: %v\nExpected: %v", tc.err, rec.Code, tc.expectedCode)
			}
		})
	}

}
