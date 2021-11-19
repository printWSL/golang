package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		data := gin.H{
			"name":    "小王子",
			"message": "hello world!",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/another_json", func(c *gin.Context) {
		type msg struct {
			Name    string
			Message string
			Age     int
		}

		data := msg{
			Name:    "小王子",
			Message: "Hello golang!",
			Age:     18,
		}

		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
