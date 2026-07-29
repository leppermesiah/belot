package main

import (
	"flag"
	"log"
	"net/http"

	"belot/server"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	webDir := flag.String("web", "./web", "path to static web client files")
	flag.Parse()

	hub := server.NewHub()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*webDir)))
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWS(hub, w, r)
	})

	log.Printf("belot server listening on %s (serving %s)\n", *addr, *webDir)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
