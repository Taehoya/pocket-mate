package entities

type Country struct {
	ID       int    `json:"id" example:"1"`
	Code     string `json:"code" example:"CA"`
	Name     string `json:"name" example:"Canada"`
	Currency string `json:"currency" example:"$"`
}
