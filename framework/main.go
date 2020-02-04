package main

import (
	"net/http"
	"practice/framework/core"
)

func main() {
	r := core.New()
	r.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	r.GET("/hello", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s \n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9000")
}
