package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yushchuk6/multidraw/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		ID:   string(p),
		Conn: conn,
		Pool: pool,
	}

	pool.Clients[client] = true

	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
