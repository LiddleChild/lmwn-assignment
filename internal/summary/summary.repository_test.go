package summary

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

func TestGetCovidCases(t *testing.T) {
	testCases := []struct {
		name          string
		invalidURL    bool
		code          int
		body          string
		expectedError bool
	}{
		{
			name:       "response OK",
			invalidURL: false,
			code:       http.StatusOK,
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
			expectedError: false,
		},
		{
			name:          "response OK, parse fail",
			invalidURL:    false,
			code:          http.StatusOK,
			body:          "{",
			expectedError: true,
		},
		{
			name:          "response not found",
			invalidURL:    false,
			code:          http.StatusNotFound,
			expectedError: true,
		},
		{
			name:          "invalid url",
			invalidURL:    true,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
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
		if tc.invalidURL {
			url = ""
		}

		err := repo.GetCovidCases(&cases, url)

		if err != nil != tc.expectedError {
			t.Errorf("\nResponse code: %v\nResponse body: '%v'\nExpected error: %v\nActual error: %v", tc.code, tc.body, tc.expectedError, err)
		}
	}
}
