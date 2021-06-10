package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"tic-tac-toe/websocket"
)

func main() {
	fmt.Println("Starting webserver...")

	addr := flag.String("addr", ":8080", "http service address")

	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/", websocket.ServeIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.Serve(hub, w, r)
	})
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
