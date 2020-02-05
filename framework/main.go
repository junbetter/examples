package main

import (
	"examples/framework/core"
	"log"
	"net/http"
	"time"
)

func onlyForV2() core.HandlerFunc {
	return func(c *core.Context) {
		t := time.Now()
		// if a server error occurred
		c.Fail(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.R.RequestURI, time.Since(t))

	}
}

func main() {
	r := core.New()
	r.Use(core.Logger())
	r.GET("/index", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *core.Context) {
			c.HTML(http.StatusOK, "<h1>hello world</h1>")
		})

		v1.GET("/hello/:name", func(c *core.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *core.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *core.Context) {
			c.JSON(http.StatusOK, core.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.GET("/assets/*filepath", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9000")
}
