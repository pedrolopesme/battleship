package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	domain "github.com/pedrolopesme/battleship/internal/domain"
)

// TODO make it all internal
type Client struct {
	Conn   *websocket.Conn
	Pool   *Pool
	Player *domain.Player
}

func (c *Client) Listen() {
	// As long as we cannot listen to websocket client
	// we can assume that the client was disconnected
	defer c.Disconnect()

	for {
		_, payload, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err) // TODO replace by proper logging
			return
		}

		log.Printf("Message received: %+v\n", payload) // TODO: replace by proper logging
	}
}

// Disconnect removes client from pool
func (c *Client) Disconnect() {
	c.Pool.unregister <- c
	c.Conn.Close()
}

func (c Client) sendMessage(socketMessage WsMessage) {
	message, _ := json.Marshal(socketMessage)
	if err := c.Conn.WriteMessage(1, message); err != nil {
		fmt.Println(err.Error())
		return
	}
}
