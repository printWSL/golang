package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 创建一个默认路由

	//加载静态文件/static表示在tmpl中引用的目录./static加载目录所在位置
	r.Static("/xxx", "./static")

	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/**/*") // templates目录下的所有文件夹和所有文件

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts/index.tmpl",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})

	r.Run(":9090")
}
