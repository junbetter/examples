package core

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type IEngine interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Run(addr string) (err error)
}

type IRouterGroup interface {
	addRoute(method string, pattern string, handler HandlerFunc)
	GET(pattern string, handler HandlerFunc)
	POST(pattern string, handler HandlerFunc)
	Group(prefix string) *RouterGroup
	Use(middlewares ...HandlerFunc)
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share Engine instance
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

// New is the constructor of Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Default use Logger() & Recovery middleware
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

// Group is defined to create a new RouteGroup
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	e := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: e,
	}
	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
