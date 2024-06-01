// study_request/4.文件上传.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("users", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fileHeader.Filename) // 文件名
		fmt.Println(fileHeader.Size)     // 文件大小，单位是字节

		//file, _ := fileHeader.Open()
		//byteData, _ := io.ReadAll(file)
		//
		//err = os.WriteFile("xxx.jpg", byteData, 0666)
		//fmt.Println(err)
		err = c.SaveUploadedFile(fileHeader, "uploads/xxx/yyy/"+fileHeader.Filename)
		fmt.Println(err)
	})
	r.Run(":8080")
}
