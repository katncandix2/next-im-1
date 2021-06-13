package main

import (
	"flag"
	"net/http"
	"next-im/pkg/chat/handler"
	log "next-im/pkg/log"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

	flag.Parse()

	hub := handler.NewHub()
	go hub.Run()

	http.HandleFunc("/", handler.ServeHomeHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWsHandler(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err == nil {
		log.GetLogger().Info("Listen Server: ", addr)
	}

}
