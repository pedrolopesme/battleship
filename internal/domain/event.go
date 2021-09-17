package domain

const (
	EVENT_NEW_GAME = "new_game"
)

type Event struct {
	EventType string `json:"type"`
	Message   string `json:"message"`
}
