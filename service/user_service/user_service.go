package user_service

import (
	db "github.com/Han-Ya-Jun/gin_scaffold/models/mongo"
)

/*
* @Author:hanyajun
* @Date:2019/7/11 23:53
* @Name:user_service
* @Function:user_service
 */

type UserService struct {
	User *db.User
}

func (u *UserService) Add() error {
	return db.InsertUser(u.User)
}
