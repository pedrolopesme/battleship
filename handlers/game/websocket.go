package game

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

type GameWebsocket struct {
	upgrader websocket.Upgrader
	socket   *websocket.Conn
}

func NewGameWebsocket() *GameWebsocket {
	return &GameWebsocket{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
		},
	}
}

// Run fires up the main websocket that controls a game battle
func (gws *GameWebsocket) Run(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	conn, err := gws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	gws.socket = conn

	gws.listen()
}

// Listing is called whenever an event happens
func (gws *GameWebsocket) listen() {
	if gws.socket == nil {
		return
	}

	for {
		messageType, message, err := gws.socket.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(messageType, message)
	}
}
