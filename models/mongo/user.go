package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/7/11 23:05
* @Name:models
* @Function: hub
 */

const (
	collectionUser = "user"
)

type User struct {
	UserId     string    `json:"user_id" bson:"_id"`
	UserName   string    `json:"user_name" bson:"user_name"`
	CreateTime time.Time `json:"-" bson:"create_time"`
}

type PageBean struct {
	PageCount  int         `json:"page_count"`
	PageSize   int         `json:"page_size"`
	PageNo     int         `json:"page_no"`
	TotalCount int         `json:"total_count"`
	List       interface{} `json:"list"`
}

// batch insert users
func InsertUsers(users []*User) error {
	docs := make([]interface{}, len(users))
	for i, u := range users {
		u.CreateTime = time.Now()
		docs[i] = u
	}
	return Inserts(db, collectionUser, docs)
}

// Find users by uid list
func FindUsers(userIdList []string) ([]*User, error) {
	var users []*User
	return users, FindAll(db, collectionUser, bson.M{"_id": map[string]interface{}{"$in": userIdList}}, nil, &users)
}

// insert user
func InsertUser(user *User) error {
	return Insert(db, collectionUser, user)
}

// update user
func UpdateUser(user *User) error {
	return Update(db, collectionUser, bson.M{"_id": user.UserId}, user)
}

// delete user
func DeleteUser(user *User) error {
	return Remove(db, collectionUser, bson.M{"_id": user.UserId})
}

// find user with page
func FindUserPage(pageSize, currentPage int) (PageBean, error) {
	pageBean := PageBean{
		PageSize: pageSize,
		PageNo:   currentPage,
	}
	count, err := Count(db, collectionUser, nil)
	if err != nil {
		return pageBean, err
	}
	var userList []*User
	err = FindPage(db, collectionUser, pageBean.PageNo, pageBean.PageSize, nil, nil, &userList)
	if err != nil {
		return pageBean, err
	}
	pageBean.PageCount = count / pageSize
	if count%pageBean.PageSize > 0 {
		pageBean.PageCount = pageBean.PageCount + 1
	}
	return pageBean, nil
}
