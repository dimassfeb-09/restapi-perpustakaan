package domain

type Officer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
