package main

import (
	"github.com/freglyc/tsuro/server"
	"net/http"
)

func main() {
	hub := server.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(response http.ResponseWriter, request *http.Request) {
		server.ServeWs(hub, response, request)
	})
}
