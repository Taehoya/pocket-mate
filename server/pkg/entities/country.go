package entities

type Country struct {
	id       int
	code     string
	name     string
	currency string
}

func NewCountry(id int, code string, name string, currency string) *Country {
	return &Country{
		id:       id,
		code:     code,
		name:     name,
		currency: currency,
	}
}
