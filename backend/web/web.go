package web

import (
	"drawAtree/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebServer(addr string) {
	router := gin.Default()
	// 启用 CORS 中间件，允许来自所有来源的请求
	router.Use(cors.Default())

	// 定义路由和处理函数
	router.POST("/api/generateTree", handler.GenerateTreeHandler)

	// 启动服务器
	router.Run(addr)
}
