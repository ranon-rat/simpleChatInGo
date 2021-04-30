package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ranon-rat/simpleChatInGo/src/stuff"
)

func ReceiveMessages(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	ws, err := stuff.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("new connection on", name)
	defer ws.Close()
	Message[name] = make(chan stuff.Messages)
	// some stuff for do this work
	// lo que estoy haciendo aqui es checar si el canal que se a ingresado , en el caso de que no exista , se cierra

	Conn, exist := Clients[name]
	if !exist {
		Conn = map[*websocket.Conn]bool{
			ws: true,
		}
	}

	Clients[name] = Conn
	Clients[name][ws] = true

	go SendMessages(name) // send the messages
	for {                 // receive the messsages
		var msg stuff.Messages

		// parse the message
		if err := ws.ReadJSON(&msg); err != nil { // if the websocket is closed
			// if the websocket is closed
			log.Println(err)
			delete(Clients[name], ws)
			// if no one is in the channel,the channel and the message is deleted
			if len(Clients[name]) == 0 {
				log.Println(len(Clients[name]))
				delete(Clients, name)
				delete(Message, name)
			}
			return

		}
		log.Println(msg)
		Message[name] <- msg
	}

}
