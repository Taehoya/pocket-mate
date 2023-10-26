package dto

import "github.com/Taehoya/pocket-mate/pkg/entities"

type CountryResponseDTO struct {
	Code     string `json:"code" example:"CA"`
	Name     string `json:"name" example:"Canada"`
	Currency string `json:"currency" example:"$"`
}

func CreateCountryResponse(country *entities.Country) *CountryResponseDTO {
	return &CountryResponseDTO{
		Code:     country.Code(),
		Name:     country.Name(),
		Currency: country.Currency(),
	}
}

func CreateCountryResponseList(countries []*entities.Country) []*CountryResponseDTO {
	countryResp := make([]*CountryResponseDTO, 0)

	for _, c := range countries {
		country := CreateCountryResponse(c)
		countryResp = append(countryResp, country)
	}

	return countryResp
}
