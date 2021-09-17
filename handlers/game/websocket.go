package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pedrolopesme/battleship/internal/domain"
	"github.com/pedrolopesme/battleship/internal/ports"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

type GameWebsocket struct {
	upgrader    websocket.Upgrader
	socket      *websocket.Conn
	gameService ports.GameService
}

func NewGameWebsocket(gameService ports.GameService) *GameWebsocket {
	return &GameWebsocket{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
		},
		gameService: gameService,
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

	gws.read()
}

// Listing is called whenever an event happens
func (gws *GameWebsocket) read() {
	if gws.socket == nil {
		return
	}

	for {
		messageType, message, err := gws.socket.ReadMessage()

		if err != nil {
			fmt.Println(err) // TODO replace by a proper logger
			return
		}

		fmt.Println(messageType, string(message))
		response := proxyEvent(gws.gameService, message)
		gws.write(response)
	}
}

func (gws *GameWebsocket) write(message string) {
	if err := gws.socket.WriteMessage(1, []byte(message)); err != nil {
		fmt.Println(err.Error())
		return
	}
}

// proxies websocket events sent by clients to game service funcs
// TODO should it be here?
func proxyEvent(gameService ports.GameService, message []byte) string {
	if len(message) == 0 {
		return ""
	}

	event := domain.Event{}
	if err := json.Unmarshal([]byte(message), &event); err != nil {
		fmt.Println("ERROR", err) // TODO replace by a proper logger
		return ""
	}

	if event.EventType == domain.EVENT_NEW_GAME {
		return createGame(gameService)
	} else {
		return ""
	}
}

// creates a game and return its json representation
func createGame(gameService ports.GameService) string {
	game, err := gameService.Create()
	if err != nil {
		fmt.Println("ERROR", err) // TODO replace by a proper logger
		return ""
	}

	gameJson, err := json.Marshal(game)
	if err != nil {
		fmt.Println("ERROR", err) // TODO replace by a proper logger
		return ""
	}

	return string(gameJson)
}
