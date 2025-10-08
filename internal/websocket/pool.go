package websocket

import (
	"fmt"
)

type Pool struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

func newPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Connection Pool Size: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined!"})
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Connection Pool Size: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected!"})
			}

		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients")
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(message)
			}
		}

	}
}
