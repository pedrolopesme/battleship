package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pedrolopesme/battleship/internal/domain"
	gameDomain "github.com/pedrolopesme/battleship/internal/domain"
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
		if err := gws.proxyEvent(message); err != nil {
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
func (gws *GameWebsocket) proxyEvent(message []byte) error {
	if len(message) == 0 {
		return errors.New("no message passed")
	}

	event := domain.Event{}
	if err := json.Unmarshal([]byte(message), &event); err != nil {
		return err
	}

	if event.EventType == domain.EVENT_NEW_MATCH {
		return gws.createMatch()
	} else if event.EventType == domain.EVENT_ENTER_LOBBY {
		return gws.enterLobby(event.Message)
	}

	return errors.New("no event type mached")
}

// enterLobby assigns a player to a waiting room
func (gws *GameWebsocket) enterLobby(payload string) error {
	player := gameDomain.Player{}
	if err := json.Unmarshal([]byte(payload), &player); err != nil {
		return err
	}

	return nil
}

// createMatch creates a game and send its json representation to client
func (gws *GameWebsocket) createMatch() error {
	game, err := gws.gameService.Create()
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
