package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// TODO make it all internal
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, payload, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err) // TODO replace by proper logging
			return
		}

		message := WsMessage{MessageType: messageType, MessageBody: string(payload)}
		c.Pool.broadcast <- message
		log.Printf("Message received: %+v\n", message) // TODO: replace by proper logging
	}
}

func (c Client) sendMessage(socketMessage WsMessage) {
	message, _ := json.Marshal(socketMessage)
	if err := c.Conn.WriteMessage(1, message); err != nil {
		fmt.Println(err.Error())
		return
	}
}
