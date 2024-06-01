// /3.中间件.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home")
	fmt.Println(c.Get("GM1"))
	fmt.Println(c.Get("GM2"))
	_user, ok := c.Get("user")
	if ok {
		user, ok := _user.(UserInfo)
		if ok {
			fmt.Println(user.Name)
			panic("xxx")
		}
	}
	c.String(200, "Home")
}
func M1(c *gin.Context) {
	fmt.Println("M1 请求部分")
	c.Abort()
	fmt.Println("M1 响应部分")
}
func M2(c *gin.Context) {
	fmt.Println("M2 请求部分")
	c.Next()
	fmt.Println("M2 响应部分")
}

type UserInfo struct {
	Name string
}

func GM1(c *gin.Context) {
	fmt.Println("GM1 请求部分")
	var user = UserInfo{Name: "枫枫"}
	c.Set("GM1", "GM1")
	c.Set("user", user)
	c.Next()
	fmt.Println("GM1 响应部分")
}

func GM2(c *gin.Context) {
	fmt.Println("GM2 请求部分")
	c.Set("GM2", "GM2")
	c.Next()
	fmt.Println("GM2 响应部分")
	fmt.Println(c.Get("GM2"))
}

func AuthMiddleware(c *gin.Context) {

}

func main() {
	r := gin.Default()
	g := r.Group("api")
	g.Use(GM1, GM2)
	g.GET("users", Home)
	r.Run(":8080")
}
