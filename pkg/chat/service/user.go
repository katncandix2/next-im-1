package service

import (
	"next-im/pkg/chat/dao"
	"next-im/pkg/chat/model"
)

type UserService struct {
	UserDao *dao.UserDao
}

func (userService *UserService) AddFriends() bool {
	//check friends exist

	//add friends
	return false
}

func (userService *UserService) CrateGroup() {
}

func (userService *UserService) JoinGroup() {
}

func (userService *UserService) GetFriendsListById(id int) []model.User {
	return userService.UserDao.GetFriendsByUid(id)
}
