// study_response/3.响应文件.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream") // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Disposition", "attachment; filename=3.响应文件.go")
		c.File("3.响应文件.go")
	})
	r.Run(":8080")
}
