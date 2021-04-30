package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/ranon-rat/simpleChatInGo/src/stuff"
)

var (
	//Clients[name of channel]websocket
	Clients = make(map[string]map[*websocket.Conn]bool)
	//Message[Name of channel] message json structure
	Message = make(map[string]chan stuff.Messages)
)
