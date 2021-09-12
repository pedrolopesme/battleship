package domain

type Ship struct {
	Destroyed bool     `json:"destroyed"`
	Location  Location `json:"location"`
}

type Location struct {
	Col uint `json:"column"`
	Row uint `json:"row"`
}
