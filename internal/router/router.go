package router

import (
	"net/http"
)

// HandlerFunc defines a function type for handling HTTP requests.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Middleware defines a function type for middleware.
type Middleware func(HandlerFunc) HandlerFunc

// Router is responsible for managing routes and middlewares.
type Router struct {
	routes      map[string]map[string]HandlerFunc
	middlewares []Middleware
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{
		routes:      make(map[string]map[string]HandlerFunc),
		middlewares: []Middleware{},
	}
}

// Use adds a middleware to the Router.
func (r *Router) Use(mw Middleware) *Router {
	r.middlewares = append(r.middlewares, mw)
	return r
}

// AddRoute adds a route to the Router.
func (r *Router) AddRoute(method, pattern string, handler HandlerFunc) *Router {
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][pattern] = handler
	return r
}

// Get adds a GET route to the Router.
func (r *Router) Get(pattern string, handler HandlerFunc) *Router {
	return r.AddRoute(http.MethodGet, pattern, handler)
}

// Post adds a POST route to the Router.
func (r *Router) Post(pattern string, handler HandlerFunc) *Router {
	return r.AddRoute(http.MethodPost, pattern, handler)
}

// applyMiddlewares applies the registered middlewares to the handler.
func (r *Router) applyMiddlewares(handler HandlerFunc) HandlerFunc {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	return handler
}

// ServeHTTP handles incoming HTTP requests and matches them to the correct route.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, exists := r.routes[req.Method]; exists {
		if handler, ok := handlers[req.URL.Path]; ok {
			// Apply middlewares to the handler
			handler = r.applyMiddlewares(handler)
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
