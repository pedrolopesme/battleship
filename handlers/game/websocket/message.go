package websocket

const (
	EVENT_ENTER_LOBBY = iota
	EVENT_NEW_MATCH
)

type WsMessage struct {
	MessageType int    `json:"type"`
	MessageBody string `json:"body"`
}
