package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Message from socket %v", string(message))
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	dir := "./static"

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	http.HandleFunc("/ws", websocketHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
