// study_request/3.form表单.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("users", func(c *gin.Context) {
		name := c.PostForm("name")
		age, ok := c.GetPostForm("age")
		fmt.Println(name)
		fmt.Println(age, ok)
	})
	r.Run(":8080")
}
