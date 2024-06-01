// study_binding/2.参数校验.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("json", func(c *gin.Context) {
		type User struct {
			IPList []string `json:"ipList" binding:"ip"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
