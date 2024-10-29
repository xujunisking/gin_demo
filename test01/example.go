package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() //监听并在 0.0.0.0:8080上启动服务
}

//# 运行 example.go 并且在浏览器中访问 HOST_IP:8080/ping
//$ go run example.go
