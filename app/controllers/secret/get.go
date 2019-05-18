package secret

import (
	"fastvault/app/controllers"
	"github.com/valyala/fasthttp"
	"net/http"
)

const (
	TOKEN_HEADER_KEY = "X-Application-Token"
)

func Get(ctx *fasthttp.RequestCtx) {
	token := ctx.Request.Header.Peek(TOKEN_HEADER_KEY)
	if token == nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		controllers.ResponseApplicationError(
			ctx,
			controllers.ApplicationErrorResponse{Error: "Token is missing"})
		return
	}

	filename := filename(token)
	content, err := fileService.ReadSecretFile(string(filename))
	if err != nil {
		ctx.SetStatusCode(http.StatusNotFound)
		controllers.ResponseApplicationError(
			ctx,
			controllers.ApplicationErrorResponse{Error: "Token is not found"})
		return
	}

	controllers.ResponseSecretReadResponse(ctx, content)
}

//24d648245969af6c97dc5ce5b93c34a39e87e411ab35b50a63258240a7dd35efe85d649819eb3bf54c7d2b9c10344ce6bd21b015106a14cf59ad048e9c67cddc.secret