package app

import "github.com/valyala/fasthttp"

type Router struct {
	route map[string]map[string]func(ctx *fasthttp.RequestCtx)
}

func (rt *Router) GetModule() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		path := ctx.Request.URI().Path()
		method := ctx.Request.Header.Method()
		rt.route[string(method)][string(path)](ctx)
	}
}

func NewRouter() *Router {
	route := make(map[string]map[string]func(ctx *fasthttp.RequestCtx))
	return &Router{
		route: route,
	}
}

func (rt *Router) Get(method, path string) func(ctx *fasthttp.RequestCtx) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	return rt.route[method][path]
}

func (rt *Router) LinkHandler(method, path string, handler Handler) {
	if rt.route[method] == nil {
		rt.route[method] = make(map[string]func(ctx *fasthttp.RequestCtx))
	}

	rt.route[method][path] = handler.Handle
}

func (rt *Router) LinkFunc(method, path string, f func(ctx *fasthttp.RequestCtx)) {
	if rt.route[method] == nil {
		rt.route[method] = make(map[string]func(ctx *fasthttp.RequestCtx))
	}

	rt.route[method][path] = f
}
