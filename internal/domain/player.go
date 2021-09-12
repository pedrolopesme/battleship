package domain

type Player struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Ships []Ship `json:"ships"`
}
