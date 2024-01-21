package repository

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GraphZC/lmwn-jr-2024/internal/model"
)

type CovidRepository interface {
	GetCovidCases() ([]model.CovidCase, error)
}

type covidApiRepository struct {
	ApiUrl string
}

func NewCovidApiRepository(apiUrl string) CovidRepository {
	return &covidApiRepository{
		ApiUrl: apiUrl,
	}
}

func (c *covidApiRepository) GetCovidCases() ([]model.CovidCase, error) {
	// Make an request to API
	res, err := http.Get(c.ApiUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Convert response body to byte array
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Serialize byte array to struct
	var covidResponse model.CovidCaseResponse
	json.Unmarshal(body, &covidResponse)

	return covidResponse.Data, nil
}