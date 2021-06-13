package main

import (
	"flag"
	"fmt"
	"net/http"

	"next-im/pkg/chat/handler"
	"next-im/pkg/log"
	"next-im/pkg/oauth"
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

	//gitbub 回调地址
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		meta := oauth.GetUserMeta(code)
		fmt.Println(meta)
		str := "success"
		w.Write([]byte(str))
	})

	err := http.ListenAndServe(*addr, nil)
	if err == nil {
		log.GetLogger().Info("Listen Server: ", addr)
	}
}
