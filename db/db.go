package db

import (
	"fmt"
	"gin-web/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func init() {

}

//全局查询的db
var DB = GetDB()

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Mysql_user,
		conf.Mysql_password,
		conf.Mysql_host,
		conf.Mysql_db_name))
	if err != nil {
		log.Println(err)
	}

	db.DB().SetConnMaxLifetime(5 * time.Second)
	//设置连接池空闲连接数
	db.DB().SetMaxIdleConns(10)
	//设置连接池最大连接数
	db.DB().SetMaxOpenConns(60)

	//设置表名禁用复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	return db
}

//需要继承的模型类
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}
