package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IhsanAlhakim/chat-backend/internal/websocket"
)

func main() {
	PORT := ":8080"

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWebsocket(pool, w, r)
	})

	fmt.Println("React Go ChatApp v1")
	fmt.Println("Server started at Port" + PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Println("Shutting down server...")
	}
}
