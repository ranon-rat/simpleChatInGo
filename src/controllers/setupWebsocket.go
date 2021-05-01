package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ranon-rat/simpleChatInGo/src/stuff"
)

func SetupWebsocket(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	ws, err := stuff.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("new connection on", name)
	defer ws.Close()

	// some stuff for do this work
	// lo que estoy haciendo aqui es checar si el canal que se a ingresado , en el caso de que no exista , se cierra

	if _, exist := Clients[name]; !exist {
		Message[name] = make(chan stuff.Messages)
		Clients[name] = make(map[*websocket.Conn]bool)
	}
	Clients[name][ws] = true

	go SendMessages(name) // send the messages
	ReceiveMessages(ws, name)

}
