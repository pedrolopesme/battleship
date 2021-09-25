package domain

const (
	EVENT_NEW_MATCH = "new_match"
)

type Event struct {
	EventType string `json:"type"`
	Message   string `json:"message"`
}
