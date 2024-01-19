package summary

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockService struct {
	getSummary func(*Summary) error
}

func (h *mockService) GetSummary(summary *Summary) error {
	return h.getSummary(summary)
}

func TestGetSummary(t *testing.T) {
	testCases := []struct {
		name         string
		err          error
		expectedCode int
	}{
		{
			name:         "OK",
			err:          nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Internal Server Error",
			err:          errors.New("Error occured while retrieving data from server."),
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := mockService{
				getSummary: func(s *Summary) error {
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
