package domain

type Game struct {
	ID    string `json:"id"`
	State uint   `json:"state"`
	Board Board  `json:"board"`
}
