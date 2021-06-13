package dao

import "next-im/pkg/chat/model"

type UserDao struct {

}

func (userDao *UserDao) GetUserInfoByUid() {

}

func (userDao *UserDao) GetGroupsByUid() {

}

func (userDao *UserDao) GetFriendsByUid(id int) []model.User {
	return nil
}

func (userDao *UserDao) AddFriends() bool {
	return true
}
