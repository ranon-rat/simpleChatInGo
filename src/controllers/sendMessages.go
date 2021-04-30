package controllers

func SendMessages(name string) {
	for {
		msg := <-Message[name] // wait for the message

		for Client := range Clients[name] {

			if err := Client.WriteJSON(msg); err != nil { // if the socket is closed

				delete(Clients[name], Client)

				// if no one is in the channel,the channel and the message is deleted

				if len(Clients[name]) == 0 {
					delete(Clients, name)
					delete(Message, name)
				}
				return
			}
		}
	}

}
