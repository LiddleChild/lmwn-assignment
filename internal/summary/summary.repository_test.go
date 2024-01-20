package summary

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LiddleChild/covid-stat/apperror"
	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

func TestGetCovidCases(t *testing.T) {
	testCases := []struct {
		name        string
		validUrl    bool
		code        int
		body        string
		expectedErr *apperror.AppError
	}{
		{
			name:     "response ok",
			validUrl: false,
			code:     http.StatusOK,
			body: `{
				"Data": [
					{
						"Age": 51,
						"Province": "Phrae"
        	},
					{
            "Age": 52,
            "Province": "Chumphon"
					}
				]
			}`,
			expectedErr: nil,
		},
		{
			name:        "response ok decode fail",
			validUrl:    false,
			code:        http.StatusOK,
			body:        "{",
			expectedErr: apperror.DecodeError,
		},
		{
			name:        "response not found",
			validUrl:    false,
			code:        http.StatusNotFound,
			expectedErr: apperror.ResponseError,
		},
		{
			name:        "service unavailable",
			validUrl:    true,
			expectedErr: apperror.ServiceUnavailable,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.code)
				_, err := w.Write([]byte(tc.body))
				if err != nil {
					fmt.Println(err)
				}
			}))
			defer testServer.Close()

			repo := NewRepository()

			cases := []covid_case.CovidCase{}

			url := testServer.URL
			if tc.validUrl {
				url = ""
			}

			err := repo.GetCovidCases(&cases, url)

			if err != tc.expectedErr {
				t.Errorf("\nResponse code: %v\nResponse body: '%v'\nExpected error: %v\nActual error: %v\n", tc.code, tc.body, tc.expectedErr, err)
			}
		})
	}
}
