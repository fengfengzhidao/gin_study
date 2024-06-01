// study_response/4.静态文件.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("st", "static")
	r.StaticFile("abc", "static/abc.txt")
	r.GET("abc", func(c *gin.Context) {
		c.String(200, "你好")
	})
	r.Run(":8080")
}
