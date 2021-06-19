package model

import "gin-web/db"

type User struct {
	db.Model
	Username         string  `json:"username"`
	Password         string  `json:"password"`
}

//根据用户名查询用户
func GetUserByUserName(userName string) User {
	var user User
	db.DB.First(&user, "username=?", userName)
	return user
}


