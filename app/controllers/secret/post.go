package secret

import (
	"fastvault/app/controllers"
	"github.com/valyala/fasthttp"
	"net/http"
)

/*
	POST /secret
*/

func Post(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	token, err := secretService.CreateSecret(body)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		controllers.ResponseApplicationError(
			ctx,
			controllers.ApplicationErrorResponse{Error: err.Error()})
		return
	}

	filename := filename([]byte(token))
	err = fileService.CreateSecretFile(filename, body)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		controllers.ResponseApplicationError(
			ctx,
			controllers.ApplicationErrorResponse{Error: err.Error()})
		return
	}

	controllers.ResponseCreateSecretResponse(ctx, controllers.SecretCreatedResponse{Token: token})
}
