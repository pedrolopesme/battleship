package domain

const (
	EVENT_ENTER_LOBBY = "enter_lobby"
	EVENT_NEW_MATCH   = "new_match"
)

type Event struct {
	EventType string `json:"type"`
	Message   string `json:"message"`
}
