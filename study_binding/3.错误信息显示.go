package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		name := field.Tag.Get("json")
		return fmt.Sprintf("%s---%s", name, label)
	})
}

/*
{
	"name": "name参数必填",
}
*/

func ValidateErr(err error) any {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	var m = map[string]any{}
	for _, e := range errs {
		msg := e.Translate(trans)
		_list := strings.Split(msg, "---")
		m[_list[0]] = _list[1]
	}
	return m
}

type User struct {
	Name  string `json:"name" binding:"required" label:"用户名"`
	Email string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()
	// 注册路由
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			// 参数验证失败
			c.JSON(200, map[string]any{
				"code": 7,
				"msg":  "验证错误",
				"data": ValidateErr(err),
			})
			return
		}

		// 参数验证成功
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s! Your email is %s.", user.Name, user.Email),
		})
	})

	// 启动HTTP服务器
	r.Run()
}
