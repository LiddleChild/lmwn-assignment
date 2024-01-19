package summary

import (
	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

type Service interface {
	GetSummary(*[]covid_case.CovidCase) error
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo,
	}
}

func (s *serviceImpl) GetSummary(result *[]covid_case.CovidCase) error {
	err := s.repo.GetCovidCases(result)
	if err != nil {
		return err
	}

	return nil
}
