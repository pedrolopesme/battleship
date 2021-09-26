package websocket

type WsMessage struct {
	MessageType int    `json:"type"`
	MessageBody string `json:"body"`
}
