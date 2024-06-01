// study_request/1.查询参数.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "25")
		keyList := c.QueryArray("key")
		fmt.Println(name, age, keyList)
	})
	r.Run(":8080")
}
