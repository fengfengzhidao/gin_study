package main

import (
	"gin_study/study_response/res"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": gin.H{}})
		res.OkWithMsg(c, "登陆成功")
	})
	r.GET("/users", func(c *gin.Context) {
		// c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": gin.H{}})
		res.OkWithData(c, map[string]any{
			"name": "枫枫",
		})
	})
	r.POST("/users", func(c *gin.Context) {
		// c.JSON(200, gin.H{"code": 0, "msg": "成功", "data": gin.H{}})
		res.FailWithMsg(c, "参数错误")
	})

	r.Run(":8080")
}
