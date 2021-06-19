package main

import (
	"fmt"
	"gin-web/conf"
	"gin-web/db"
	"gin-web/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,access-token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	r := view.IniteRouter()
	//跨域
	r.Use(Cors())

	defer db.DB.Close()

	err := r.Run(fmt.Sprintf(":%s", conf.Server_port))
	if err != nil {
		log.Fatalf("return err: #{err}")
	}
}
