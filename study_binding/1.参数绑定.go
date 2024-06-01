// study_binding/1.参数绑定.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBindQuery(&user)
		fmt.Println(user, err)
	})

	r.GET("users/:id/:name", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			ID   int    `uri:"id"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		fmt.Println(user, err)
	})

	r.POST("form", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBind(&user)
		fmt.Println(user, err)
	})

	r.POST("json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println(user, err)
	})
	r.POST("header", func(c *gin.Context) {
		type User struct {
			Name        string `header:"Name"`
			Age         int    `header:"Age"`
			UserAgent   string `header:"User-Agent"`
			ContentType string `header:"Content-Type"`
		}
		var user User
		err := c.ShouldBindHeader(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
