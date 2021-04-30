package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleChatInGo/src/controllers"
)

func SetupRouter() error {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/home.html")
	})
	r.HandleFunc(`/channel/{name:[\w\W]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/channel.html")
	})
	r.HandleFunc(`/public/{file:[\w\W\/]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	r.HandleFunc(`/ws/{name:[\w\W]}`, controllers.ReceiveMessages)
	log.Println("runnin on http://localhost:8080")
	return http.ListenAndServe(":8080", r)
}
