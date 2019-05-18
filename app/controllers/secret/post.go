package secret

import (
	"fastvault/app/controllers"
	"fastvault/app/utils"
	"fmt"
	"github.com/valyala/fasthttp"
)

/*
	POST /secret
*/

func Post(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	token, err := secretService.CreateSecret(body)
	if err != nil {
		controllers.ResponseApplicationError(
			ctx,
			controllers.ApplicationErrorResponse{Error: err.Error()})
		return
	}

	filename := fmt.Sprintf("%x", utils.ToSha512([]byte(token)))
	err = fileService.CreateSecretFile(filename, body)
	if err != nil {
		controllers.ResponseApplicationError(
			ctx,
			controllers.ApplicationErrorResponse{Error: err.Error()})
		return
	}

	controllers.ResponseCreateSecretResponse(ctx, controllers.SecretCreatedResponse{Token: token})
}
