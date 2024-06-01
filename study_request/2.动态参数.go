// study_request/2.动态参数.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		fmt.Println(userID)
	})
	r.GET("users/:id/:name", func(c *gin.Context) {
		userID := c.Param("id")
		userName := c.Param("name")
		fmt.Println(userID, userName)
	})
	r.Run(":8080")
}
