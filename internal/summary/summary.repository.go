package summary

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LiddleChild/covid-stat/apperror"
	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

type Repository interface {
	GetCovidCases(*[]covid_case.CovidCase, string) *apperror.AppError
}

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetCovidCases(result *[]covid_case.CovidCase, url string) *apperror.AppError {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error occured while requesting data from covid stat server. %v\n", err.Error())
		return apperror.ServiceUnavailable
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		fmt.Printf("Covid stat server responded with code %v\n", res.StatusCode)
		return apperror.ResponseError
	}

	defer res.Body.Close()

	var casesReponse covid_case.CovidCasesResponse

	err = json.NewDecoder(res.Body).Decode(&casesReponse)
	if err != nil {
		fmt.Printf("Error occured while decoding response body. %v\n", err.Error())
		return apperror.DecodeError
	}

	*result = casesReponse.Data

	return nil
}
