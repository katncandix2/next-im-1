package main

import (
	"flag"
	"net/http"
	"time"

	"next-im/pkg/chat/constant"
	"next-im/pkg/chat/db"
	"next-im/pkg/chat/handler"
	"next-im/pkg/log"
	"next-im/pkg/oauth"

	"github.com/dgrijalva/jwt-go"
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
	http.HandleFunc("/add_friend", handler.AddFriendHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWsHandler(hub, w, r)
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/chat/front/auth.html")
	})

	//gitbub 回调地址
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		meta := oauth.GetUserMeta(code)
		claims := jwt.MapClaims{
			"user":meta,
			"exp": time.Now().Add(time.Duration(60 * 60 * 24)*time.Second).Unix(), // 过期时间，必须设置,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(oauth.SecretDevKey))
		if err!=nil{
			log.GetLogger().Errorln(err)
		}

		//todo Authorization 全局有效
		w.Header().Set("Authorization",tokenString)
		w.Write([]byte("success"))
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
