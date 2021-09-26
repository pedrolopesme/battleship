package websocket

import (
	"log"
)

type Pool struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan WsMessage
}

func NewPool() *Pool {
	return &Pool{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan WsMessage),
	}
}

func (p *Pool) BroadcastMessage(message string) {
	log.Println("Broadcasting Message:", message)
	for client, _ := range p.clients {
		log.Println("Sending message to ", client)
		client.Conn.WriteJSON(WsMessage{MessageType: 1, MessageBody: message})
	}
}

// Activate starts the pool and listen to Ws clients
func (p *Pool) Activate() {
	for {
		select {
		case client := <-p.register:
			p.clients[client] = true
			log.Println("New client onboard. Pool size:", len(p.clients))
			p.BroadcastMessage("New client joined...")
		case client := <-p.unregister:
			delete(p.clients, client)
			log.Println("Client logged out. Pool size:", len(p.clients))
			p.BroadcastMessage("Client getting out from the match")
		case message := <-p.broadcast:
			p.BroadcastMessage(message.MessageBody)
		}
	}
}
