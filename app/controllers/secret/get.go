package secret

import (
	"fastvault/app/controllers"
	"github.com/valyala/fasthttp"
	"net/http"
)

const (
	TOKEN_HEADER_KEY = "X-Application-Token"
	ACCEPT_HEADER_KEY = "Accept"
	CONTENT_TYPE_HEADER_KEY = "Content-Type"
)

func Get(ctx *fasthttp.RequestCtx) {
	token := ctx.Request.Header.Peek(TOKEN_HEADER_KEY)
	accept := ctx.Request.Header.Peek(ACCEPT_HEADER_KEY)

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

	if accept != nil {
		ctx.SetContentType(string(accept))
	}

	controllers.ResponseSecretReadResponse(ctx, content)
}