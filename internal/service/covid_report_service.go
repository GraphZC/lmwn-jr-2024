package service

import (
	"github.com/GraphZC/lmwn-jr-2024/internal/model"
	"github.com/GraphZC/lmwn-jr-2024/internal/repository"
)

var covidRepository = repository.NewCovidApiRepository("https://static.wongnai.com/devinterview/covid-cases.json")

func GetCovidCases() ([]model.CovidCase, error) {
	covidCases, err := covidRepository.GetCovidCases()
	if err != nil {
		return nil, err
	}

	return covidCases, nil
}

func SummaryAgeGroup(ageGroup *model.CovidAgeGroup, covidCase model.CovidCase) {
	if covidCase.Age == 0 {
		ageGroup.AgeGroupNA++
	} else if covidCase.Age <= 30 {
		ageGroup.AgeGroup0_30++
	} else if covidCase.Age <= 60 {
		ageGroup.AgeGroup31_60++
	} else if covidCase.Age > 60 {
		ageGroup.AgeGroup61++
	}
}

func SummaryProvince(provinceMap *map[string]int, covidCase model.CovidCase) {
	if covidCase.Province == "" {
		// If province is empty, set it to "N/A"
		(*provinceMap)["N/A"]++
	} else {
		(*provinceMap)[covidCase.Province]++
	}
}

func GetCovidSummary(covidReport []model.CovidCase) model.CovidSummary {
	// Initialize province map
	provinceMap := make(map[string]int)
	provinceMap["N/A"] = 0

	// Initialize age group
	ageGroup := model.CovidAgeGroup{
		AgeGroup0_30:  0,
		AgeGroup31_60: 0,
		AgeGroup61:    0,
		AgeGroupNA:    0,
	}

	for _, covidCase := range covidReport {
		SummaryProvince(&provinceMap, covidCase)
		SummaryAgeGroup(&ageGroup, covidCase)
	}

	covidSummary := model.CovidSummary{
		Province: provinceMap,
		AgeGroup: ageGroup,
	}

	return covidSummary
}
