package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Yushchuk6/multidraw/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket connect!")
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

	http.Handle("/", http.FileServer(http.Dir("./public")))
}

func main() {
	fmt.Println("Multidraw v0.2")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	setupRoutes()

	http.ListenAndServe(":"+port, nil)
}
