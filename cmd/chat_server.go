package main

import (
	"flag"
	"fmt"
	"net/http"

	"next-im/pkg/chat/constant"
	"next-im/pkg/chat/db"
	"next-im/pkg/chat/handler"
	log "next-im/pkg/log"
	"next-im/pkg/oauth"
)

var addr = flag.String("addr", ":8080", "http service address")

type Server struct {
	port       string
	listen     string
	configFile string
	dbEngine   int
	dbServer   string
}

func (server *Server) Init() {
	//todo: add parse config file

	//todo: default memDataAccess, support mul data
	mem := &db.MemDataAccess{}
	dbErr := mem.Init()
	if dbErr != nil {
		log.GetLogger().Error("Connect db failed")
	}
}

func (server *Server) run() {

	flag.Parse()
	hub := handler.NewHub()
	go hub.Run()
	// init db connection

	http.HandleFunc("/", handler.ServeHomeHandler)
	http.HandleFunc("/", handler.AddFriendHandler)
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

func main() {
	var server = &Server{
		dbEngine: constant.DB_ENGINE_MEM,
	}
	server.run()
}
