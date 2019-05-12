package app

import (
	"github.com/valyala/fasthttp"
	"log"
)

type Router struct {
	route map[string]map[string]func(ctx *fasthttp.RequestCtx)
}

func (rt *Router) GetModule() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Request.URI().Path())
		method := string(ctx.Request.Header.Method())
		handler := rt.Get(method, path)
		if handler != nil {
			handler(ctx)
		}else{
			ctx.SetStatusCode(404)
			ctx.Write([]byte("Not Found"))
			log.Println("[404]", method, path)
		}
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
	rt.initIfMethodBucketIsNil(method)
	rt.route[method][path] = handler.Handle
	log.Println("bind", method, path)
}

func (rt *Router) LinkFunc(method, path string, f func(ctx *fasthttp.RequestCtx)) {
	rt.initIfMethodBucketIsNil(method)
	rt.route[method][path] = f
	rt.printBindingLog(method, path)
}

func (rt *Router) printBindingLog(method string, path string) {
	log.Println("bind", method, path)
}

func (rt *Router) initIfMethodBucketIsNil(method string) {
	if rt.route[method] == nil {
		rt.route[method] = make(map[string]func(ctx *fasthttp.RequestCtx))
	}
}
