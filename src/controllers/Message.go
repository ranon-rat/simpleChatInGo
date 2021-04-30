package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleChatInGo/src/stuff"
)

func ReceiveMessages(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	ws, err := stuff.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	// some stuff for do this work
	Clients[name][ws] = true
	Message[name] = make(chan stuff.Messages)

	go SendMessages(name) // send the messages
	for {                 // receive the messsages
		var msg stuff.Messages
		// parse the message
		if err := ws.ReadJSON(&msg); err != nil { // if the websocket is closed
			// if the websocket is closed
			log.Println(err)
			delete(Clients[name], ws)
			// if no one is in the channel,the channel and the message is deleted
			if len(Clients) <= 0 {
				delete(Clients, name)
				delete(Message, name)
			}
			return

		}
		Message[name] <- msg
	}

}
