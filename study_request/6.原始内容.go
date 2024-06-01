// study_request/6.原始内容.go
package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		byteData, err := io.ReadAll(c.Request.Body)
		fmt.Println(string(byteData), err)
		// 读了body之后，body就没了
		c.Request.Body = io.NopCloser(bytes.NewReader(byteData))

		name := c.PostForm("name")
		fmt.Println(name)

	})
	r.Run(":8080")
}
