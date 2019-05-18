package secret

import (
	"fastvault/app/services"
	"github.com/valyala/fasthttp"
)

/*
	POST /secret
*/

var secretService services.SecretService = services.NewSecretService()

func Post (ctx *fasthttp.RequestCtx) {

}
