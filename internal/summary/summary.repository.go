package summary

import (
	"encoding/json"
	"net/http"

	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

type Repository interface {
	GetCovidCases(*[]covid_case.CovidCase, string) error
}

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetCovidCases(result *[]covid_case.CovidCase, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var casesReponse covid_case.CovidCasesResponse

	err = json.NewDecoder(res.Body).Decode(&casesReponse)
	if err != nil {
		return err
	}

	*result = casesReponse.Data

	return nil
}
