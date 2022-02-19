package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	ID string `json:"id"`
	X  int16  `json:"x"`
	Y  int16  `json:"y"`
}

func (c *Client) Read() {
	defer func() {
		delete(c.Pool.Clients, c)
		c.Conn.Close()
	}()

	for {
		var message Message

		err := c.Conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			return
		}
		message.ID = c.ID

		c.Pool.Broadcast <- message
	}
}
