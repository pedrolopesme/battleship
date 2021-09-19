package game

import (
	"encoding/json"
	"errors"
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
		if err := gws.proxyEvent(gws.gameService, message); err != nil {
			fmt.Println("ERROR", err) // TODO replace by a proper logger
		}
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
func (gws *GameWebsocket) proxyEvent(gameService ports.GameService, message []byte) error {
	if len(message) == 0 {
		return errors.New("no message passed")
	}

	event := domain.Event{}
	if err := json.Unmarshal([]byte(message), &event); err != nil {
		return err
	}

	if event.EventType == domain.EVENT_NEW_GAME {
		return gws.createGame(gameService)
	}

	return errors.New("no event type mached")
}

// creates a game and send its json representation to client
func (gws *GameWebsocket) createGame(gameService ports.GameService) error {
	game, err := gameService.Create()
	if err != nil {
		return err
	}

	gameJson, err := json.Marshal(game)
	if err != nil {
		return err
	}

	gws.write(string(gameJson))
	return nil
}
