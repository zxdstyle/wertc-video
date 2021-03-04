package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"server/hanlder"
	"server/middleware"
)

func main() {
	server := g.Server()

	server.Use(middleware.CORS)

	server.BindHandler("/ws", hanlder.WsHandler)

	server.Group("/", func(group *ghttp.RouterGroup) {
		// 创建聊天室
		group.POST("/room", hanlder.CreateRoom)
	})

	server.SetPort(8199)
	server.Run()
}
