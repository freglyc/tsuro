package main

import (
	"github.com/freglyc/tsuro/server"
	"log"
	"net/http"
	"os"
)

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
	//_ = os.Setenv("PORT", "8080")
	var addr = ":" + os.Getenv("PORT")
	hub := server.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(response http.ResponseWriter, request *http.Request) {
		server.ServeWs(hub, response, request)
	})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
