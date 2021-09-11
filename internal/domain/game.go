package domain

type Game struct {
	ID       string        `json:"id"`
	State    string        `json:"state"`
	Board    Board         `json:"board"`
	Settings BoardSettings `json:"board_settings"`
}

type BoardSettings struct {
	Rows uint `json:"rows"`
	Cols uint `json:"columns"`
}

type Board [][]string
