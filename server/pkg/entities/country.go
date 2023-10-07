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

func (c *Country) ID() int {
	return c.id
}

func (c *Country) Code() string {
	return c.code
}

func (c *Country) Name() string {
	return c.name
}

func (c *Country) Currency() string {
	return c.currency
}
