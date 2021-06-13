package handler

import (
	"encoding/json"
	"net/http"

	"next-im/pkg/chat/dao"
	"next-im/pkg/chat/service"
	"next-im/pkg/log"
)

type AddFriendParam struct {
	Uid       string
	FriendUid string
}

func ServeHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "pkg/chat/front/home.html")
}

func AddFriendHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var addFriendParam AddFriendParam

	err := json.NewDecoder(r.Body).Decode(&addFriendParam)

	if err != nil {
		log.GetLogger().Error("Handle Add Friend failed ", err)
	}

	userService := service.UserService{UserDao: &dao.UserDao{}}
	bool := userService.AddFriends()
	if bool {
		log.GetLogger().Info("Add friend failed")
	}

}
