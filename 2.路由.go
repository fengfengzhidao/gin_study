// /2.路由.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(c.Request.Method, path)
}

func main() {
	r := gin.Default()

	apiGroup := r.Group("api")
	apiGroup.Use()
	UserGroup(apiGroup)

	noMiddleWareGroup := r.Group("api")
	LoginGroup(noMiddleWareGroup)

	r.Run(":8080")
}

func UserGroup(r *gin.RouterGroup) {
	r.GET("users", UserView)
	r.POST("users", UserView)
	r.DELETE("users", UserView)
	r.PUT("users", UserView)
}

func LoginGroup(r *gin.RouterGroup) {
	r.GET("login", UserView)
}
