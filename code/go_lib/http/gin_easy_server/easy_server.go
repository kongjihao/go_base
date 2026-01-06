package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 定义接收JSON的结构体
type User struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	router := gin.Default()

	router.GET("/info", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	router.POST("/post", func(c *gin.Context) {
		// 服务器端使用的是 c.PostForm 来获取表单参数，
		// 这种方式只能获取 Content-Type 为 application/x-www-form-urlencoded 或 multipart/form-data的数据。
		// name := c.PostForm("name")
		// age := c.PostForm("age")

		// var user User
		// 使用 ShouldBindJSON 绑定JSON数据, 请求Content-Type 为 application/json
		// if err := c.ShouldBindJSON(&user); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// 适配两种格式
		content_type := c.GetHeader("Content-Type")
		data := make(map[string]interface{})

		if strings.Contains(content_type, "application/json") {
			var user User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			data["name"] = user.Name
			data["age"] = user.Age

		} else if strings.Contains(content_type, "application/x-www-form-urlencoded") ||
			strings.Contains(content_type, "multipart/form-data") {
			name := c.PostForm("name")
			age := c.PostForm("age")
			data["name"] = name
			data["age"] = age
		}

		c.JSON(http.StatusOK, gin.H{
			"info":   "recived request~",
			"method": "POST",
			"data":   data,
		})
	})

	router.Run(":8080")
}
