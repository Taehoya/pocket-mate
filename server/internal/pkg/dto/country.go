package dto

import "github.com/Taehoya/pocket-mate/internal/pkg/entities"

type CountryResponseDTO struct {
	Id       int    `json:"id" example:"1"`
	Code     string `json:"code" example:"CA"`
	Name     string `json:"name" example:"Canada"`
	Currency string `json:"currency" example:"$"`
}

func NewCountryResponse(country *entities.Country) *CountryResponseDTO {
	return &CountryResponseDTO{
		Id:       country.ID(),
		Code:     country.Code(),
		Name:     country.Name(),
		Currency: country.Currency(),
	}
}

func NewCountryResponseList(countries []*entities.Country) []*CountryResponseDTO {
	countryResp := make([]*CountryResponseDTO, 0)

	for _, c := range countries {
		country := NewCountryResponse(c)
		countryResp = append(countryResp, country)
	}

	return countryResp
}
