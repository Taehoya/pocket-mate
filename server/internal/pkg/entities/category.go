package entities

type Category struct {
	id   int
	name string
}

func NewCategory(id int, name string) *Category {
	return &Category{
		id:   id,
		name: name,
	}
}

func (c *Category) ID() int {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}
