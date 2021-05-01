package controllers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/ranon-rat/simpleChatInGo/src/stuff"
)

func ReceiveMessages(ws *websocket.Conn, name string) {
	for { // receive the messsages
		var msg stuff.Messages

		// parse the message
		if err := ws.ReadJSON(&msg); err != nil { // if the websocket is closed
			// if the websocket is closed
			log.Println(err)
			delete(Clients[name], ws)
			// if no one is in the channel,the channel and the message is deleted
			if len(Clients[name]) == 0 {
				delete(Clients, name) //as
				delete(Message, name)
				log.Println(Clients, Message)
			}
			return

		}
		log.Println(msg)
		Message[name] <- msg
	}

}
