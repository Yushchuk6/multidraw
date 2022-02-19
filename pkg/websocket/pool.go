package websocket

import (
	"fmt"
	"time"
)

type Pool struct {
	Clients   map[*Client]bool
	Broadcast chan Message
}

func NewPool() *Pool {
	return &Pool{
		Clients:   make(map[*Client]bool),
		Broadcast: make(chan Message, 10000),
	}
}

func (pool *Pool) Start() {
	for range time.Tick(10 * time.Millisecond) {
		pool.sendAllClients()
	}
}

func (pool *Pool) sendAllClients() {
	var messageList []Message
	for {
		select {
		case m := <-pool.Broadcast:
			messageList = append(messageList, m)
		default:
			if len(messageList) == 0 {
				return
			}
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(messageList); err != nil {
					fmt.Println(err)
					return
				}
			}
			return
		}
	}
}
