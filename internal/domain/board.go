package domain

type BoardSettings struct {
	Rows uint `json:"rows"`
	Cols uint `json:"columns"`
}

type Board struct {
	Settings BoardSettings `json:"settings"`
	Ships    []Ship        `json:"ships"`
}
