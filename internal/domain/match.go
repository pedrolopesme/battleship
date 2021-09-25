package domain

type Match struct {
	ID    string `json:"id"`
	State uint   `json:"state"`
	Board Board  `json:"board"`
}
