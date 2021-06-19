package model

import "gin-web/db"

func init() {
	db.DB.AutoMigrate(User{})
}
