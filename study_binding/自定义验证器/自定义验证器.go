package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net"
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

	v.RegisterValidation("fip", func(fl validator.FieldLevel) bool {
		fmt.Println("fl.Field(): ", fl.Field())
		fmt.Println("fl.FieldName(): ", fl.FieldName())
		fmt.Println("fl.StructFieldName(): ", fl.StructFieldName())
		fmt.Println("fl.Parent(): ", fl.Parent())
		fmt.Println("fl.Top(): ", fl.Top())
		fmt.Println("fl.Param(): ", fl.Param())

		ip, ok := fl.Field().Interface().(string)
		if ok && ip != "" {
			// 传了值就去校验是不是IP地址
			ipObj := net.ParseIP(ip)
			return ipObj != nil
		}
		return true
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
		if e.Tag() == "fip" {
			m[strings.Split(e.Field(), "---")[0]] = "该ip地址不符合要求"
			continue
		}
		m[_list[0]] = _list[1]
	}
	return m
}

func main() {
	r := gin.Default()
	// 注册路由
	r.POST("/user", func(c *gin.Context) {
		type User struct {
			Ip string `json:"ip" binding:"fip=1234" label:"ip地址"`
		}
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
		c.JSON(http.StatusOK, user)
	})

	// 启动HTTP服务器
	r.Run(":8080")
}
