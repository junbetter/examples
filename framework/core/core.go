package core

import (
	"net/http"
)

type HandlerFunc func(*Context)

type IEngine interface {
	addRoute(method string, pattern string, handler HandlerFunc)
	GET(pattern string, handler HandlerFunc)
	POST(pattern string, handler HandlerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Run(addr string) (err error)
}

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
