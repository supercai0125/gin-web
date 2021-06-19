package view

import (
	"gin-web/model"
	"github.com/gin-gonic/gin"
)


// @Summary 测试接口
// @Description get all messages
// @Accept  json
// @Produce json
// @Success 200 {string} string "success"
// @Router /hello [get]
func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"data": "hello world",
	})
	return
}

func user(c *gin.Context) {
	userName := c.Query("username")
	user := model.GetUserByUserName(userName)
	c.JSON(200, gin.H{
		"status": "success",
		"data": user,
	})
}
