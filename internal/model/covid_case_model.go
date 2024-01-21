package model

type CovidCase struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             string `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
}

type CovidCaseResponse struct {
	Data []CovidCase `json:"Data"`
}

type CovidAgeGroup struct {
	AgeGroup0_30 int `json:"0-30"`
	AgeGroup31_60 int `json:"31-60"`
	AgeGroup61 int `json:"61+"`
	AgeGroupNA int `json:"N/A"`
}

type CovidSummary struct {
	Province map[string]int `json:"Province"`
	AgeGroup CovidAgeGroup `json:"AgeGroup"`
}