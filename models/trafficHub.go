package models

type TrafficHub struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	FullName    string `json:"fullName"`
	Type        string `json:"type"`
}
