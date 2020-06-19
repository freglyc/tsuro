package main

import (
	"flag"
	"github.com/freglyc/tsuro/server"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(response, "Not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = response.Write([]byte("Tsuro Server"))
}

func main() {
	flag.Parse()
	hub := server.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(response http.ResponseWriter, request *http.Request) {
		server.ServeWs(hub, response, request)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
