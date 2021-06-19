package view

import (
	_ "gin-web/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title goweb project
// @version 1.0
// @description this is goweb server.
// @host 127.0.0.1:8080
// @BasePath /trainning
func IniteRouter() *gin.Engine {
	r := gin.New()
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//路由分组
	train := r.Group("/trainning")
	{
		train.GET("/hello", hello)
		train.GET("/user", user)
	}

	return r
}
