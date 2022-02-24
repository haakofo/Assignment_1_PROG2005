package handlers_and_structs

import "time"

type Borders struct {
	Isocodes []string `json:"borders"`
}

type Country struct {
	CountryName CountryName `json:"name"`
}
type CountryName struct {
	Common string `json:"common"`
}

type Complete_Unifinfo struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"alpha_two_code"`
	WebPages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages"`
	Maps      map[string]string `json:"maps"`
}

type Direct_diag struct {
	UniversitiesAPI string        `json:"universitiesapi"`
	CountriesSPI    string        `json:"countriesapi"`
	Version         string        `json:"name"`
	Uptime          time.Duration `json:"uptime"`
}
