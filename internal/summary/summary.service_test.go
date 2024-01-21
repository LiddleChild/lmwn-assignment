package summary

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/LiddleChild/covid-stat/apperror"
	"github.com/LiddleChild/covid-stat/config"
	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

type mockRepository struct {
	getCovidCasesFunc func(*[]covid_case.CovidCase) *apperror.AppError
}

func (r *mockRepository) GetCovidCases(result *[]covid_case.CovidCase, url string) *apperror.AppError {
	return r.getCovidCasesFunc(result)
}

func TestErrorGetSummary(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		repo := &mockRepository{
			getCovidCasesFunc: func(result *[]covid_case.CovidCase) *apperror.AppError {
				return apperror.ServiceUnavailable
			},
		}

		service := NewService(repo, &config.Config{})

		result := Summary{}
		err := service.GetSummary(&result)
		if err == nil {
			t.Errorf("This should return an error.\nActual error: %v", err)
		}
	})
}

func TestSuccessGetSummary(t *testing.T) {
	testCases := []struct {
		name     string
		testcase string
		expected Summary
	}{
		{
			name:     "success null province null age group",
			testcase: `[{ "Age": null, "Province": null }]`,
			expected: Summary{
				Province: map[string]int{
					"N/A": 1,
				},
				AgeGroup: AgeGroup{
					Young:     0,
					MiddleAge: 0,
					Elderly:   0,
					Null:      1,
				},
			},
		},
		{
			name:     "success null province",
			testcase: `[{ "Age": 1, "Province": null }]`,
			expected: Summary{
				Province: map[string]int{
					"N/A": 1,
				},
				AgeGroup: AgeGroup{
					Young:     1,
					MiddleAge: 0,
					Elderly:   0,
					Null:      0,
				},
			},
		},
		{
			name:     "success null age",
			testcase: `[{ "Age": null, "Province": "A" }]`,
			expected: Summary{
				Province: map[string]int{
					"A": 1,
				},
				AgeGroup: AgeGroup{
					Young:     0,
					MiddleAge: 0,
					Elderly:   0,
					Null:      1,
				},
			},
		},
		{
			name:     "success negative age",
			testcase: `[{ "Age": -99, "Province": "A" }]`,
			expected: Summary{
				Province: map[string]int{
					"A": 1,
				},
				AgeGroup: AgeGroup{
					Young:     0,
					MiddleAge: 0,
					Elderly:   0,
					Null:      1,
				},
			},
		},
		{
			name: "success 0-30 group",
			testcase: `[
				{ "Age":  0, "Province": "A" },
				{ "Age": 30, "Province": "B" }
			]`,
			expected: Summary{
				Province: map[string]int{
					"A": 1,
					"B": 1,
				},
				AgeGroup: AgeGroup{
					Young:     2,
					MiddleAge: 0,
					Elderly:   0,
					Null:      0,
				},
			},
		},
		{
			name: "success 31-60 group",
			testcase: `[
				{ "Age": 31, "Province": "A" },
				{ "Age": 60, "Province": "B" }
			]`,
			expected: Summary{
				Province: map[string]int{
					"A": 1,
					"B": 1,
				},
				AgeGroup: AgeGroup{
					Young:     0,
					MiddleAge: 2,
					Elderly:   0,
					Null:      0,
				},
			},
		},
		{
			name: "success 61+ group",
			testcase: `[
				{ "Age": 61, "Province": "A" }
			]`,
			expected: Summary{
				Province: map[string]int{
					"A": 1,
				},
				AgeGroup: AgeGroup{
					Young:     0,
					MiddleAge: 0,
					Elderly:   1,
					Null:      0,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockRepository{
				getCovidCasesFunc: func(result *[]covid_case.CovidCase) *apperror.AppError {
					err := json.Unmarshal([]byte(tc.testcase), result)
					if err != nil {
						fmt.Println(err.Error())
					}

					return nil
				},
			}

			service := NewService(repo, &config.Config{})

			result := Summary{}
			err := service.GetSummary(&result)
			if err != nil {
				t.Errorf("This should not return any error.\nTestcase: %v\nActual error: %v\n",
					tc.testcase,
					err,
				)
			}

			if tc.expected.AgeGroup != result.AgeGroup {
				t.Errorf("Wrong AgeGroup result.\nTestcase: %v\nResult: %v\nExpected: %v\n",
					tc.testcase,
					result,
					tc.expected,
				)
			}

			if !reflect.DeepEqual(tc.expected.Province, result.Province) {
				t.Errorf("Wrong Province result.\nTestcase: %v\nResult: %v\nExpected: %v\n",
					tc.testcase,
					result,
					tc.expected,
				)
			}
		})
	}
}
