package main

import (
	"examples/framework/core"
	"net/http"
)

func main() {
	r := core.New()
	r.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	r.POST("/login", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/hello/:name", func(c *core.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9000")
}
