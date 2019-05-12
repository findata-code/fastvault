package app

import "github.com/valyala/fasthttp"


func Start() {

}

func NewPipeline(router Router) func (ctx *fasthttp.RequestCtx) {

	router := NewRouter()

	return func(ctx *fasthttp.RequestCtx){
		router.Get(ctx.Request)
		ctx.Request.URI().Path()
	}
}
