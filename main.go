package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Multidraw start v0.1")
	setupRoutes()
	http.ListenAndServe(":"+port, nil)
}
