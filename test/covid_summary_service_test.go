package test

import (
	"testing"

	"github.com/GraphZC/lmwn-jr-2024/internal/model"
	"github.com/GraphZC/lmwn-jr-2024/internal/service"
	"github.com/stretchr/testify/assert"
)


func TestSummaryProvinceEqualNull(t *testing.T) {
	provinceMap := make(map[string]int)
	provinceMap["N/A"] = 0

	service.SummaryProvince(&provinceMap, model.CovidCase{
		Province: "",
	})

	assert.Equal(t, 1, provinceMap["N/A"])
}

func TestSummaryProvinceEqualBangkok(t *testing.T) {
	provinceMap := make(map[string]int)
	provinceMap["N/A"] = 0

	service.SummaryProvince(&provinceMap, model.CovidCase{
		Province: "Bangkok",
	})

	assert.Equal(t, 1, provinceMap["Bangkok"])
}

func TestSummaryAgeGroupEqual0(t *testing.T) {
	ageGroup := model.CovidAgeGroup{
		AgeGroup0_30:  0,
		AgeGroup31_60: 0,
		AgeGroup61:    0,
		AgeGroupNA:    0,
	}

	service.SummaryAgeGroup(&ageGroup, model.CovidCase{
		Age: 0,
	})

	assert.Equal(t, 1, ageGroup.AgeGroupNA)
}

func TestSummaryAgeGroupEqual30(t *testing.T) {
	ageGroup := model.CovidAgeGroup{
		AgeGroup0_30:  0,
		AgeGroup31_60: 0,
		AgeGroup61:    0,
		AgeGroupNA:    0,
	}

	service.SummaryAgeGroup(&ageGroup, model.CovidCase{
		Age: 30,
	})

	assert.Equal(t, 1, ageGroup.AgeGroup0_30)
}

func TestSummaryAgeGroupEqual31(t *testing.T) {
	ageGroup := model.CovidAgeGroup{
		AgeGroup0_30:  0,
		AgeGroup31_60: 0,
		AgeGroup61:    0,
		AgeGroupNA:    0,
	}

	service.SummaryAgeGroup(&ageGroup, model.CovidCase{
		Age: 31,
	})

	assert.Equal(t, 1, ageGroup.AgeGroup31_60)
}

func TestSummaryAgeGroupEqual60(t *testing.T) {
	ageGroup := model.CovidAgeGroup{
		AgeGroup0_30:  0,
		AgeGroup31_60: 0,
		AgeGroup61:    0,
		AgeGroupNA:    0,
	}

	service.SummaryAgeGroup(&ageGroup, model.CovidCase{
		Age: 60,
	})

	assert.Equal(t, 1, ageGroup.AgeGroup31_60)
}

func TestSummaryAgeGroupEqual61(t *testing.T) {
	ageGroup := model.CovidAgeGroup{
		AgeGroup0_30:  0,
		AgeGroup31_60: 0,
		AgeGroup61:    0,
		AgeGroupNA:    0,
	}

	service.SummaryAgeGroup(&ageGroup, model.CovidCase{
		Age: 61,
	})

	assert.Equal(t, 1, ageGroup.AgeGroup61)
}



