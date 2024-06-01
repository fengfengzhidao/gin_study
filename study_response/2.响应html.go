// study_response/2.响应html.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("study_response/templates/*")
	//r.LoadHTMLFiles("study_response/templates/index.html")
	r.GET("", func(c *gin.Context) {

		c.HTML(200, "index.html", map[string]any{
			"title": "枫枫知道",
		})
	})
	r.Run(":8080")
}
