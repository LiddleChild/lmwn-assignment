package summary

import (
	"github.com/LiddleChild/covid-stat/internal/covid_case"
)

type Service interface {
	GetSummary(*Summary) error
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo,
	}
}

func (s *serviceImpl) GetSummary(summary *Summary) error {
	var covidCases []covid_case.CovidCase
	err := s.repo.GetCovidCases(&covidCases)
	if err != nil {
		return err
	}

	(*summary).Province = map[string]int{}
	(*summary).AgeGroup = AgeGroup{}

	for _, c := range covidCases {
		addProvince(summary, c.Province)
		addAgeGroup(summary, c.Age)
	}

	return nil
}

func addProvince(summary *Summary, provinceName *string) {
	p := "N/A"
	if provinceName != nil {
		p = *provinceName
	}

	summary.Province[p] += 1
}

func addAgeGroup(summary *Summary, age *int) {
	if age == nil {
		summary.AgeGroup.Null += 1
	} else if *age <= 30 {
		summary.AgeGroup.Young += 1
	} else if *age <= 60 {
		summary.AgeGroup.MiddleAge += 1
	} else {
		summary.AgeGroup.Elderly += 1
	}
}
