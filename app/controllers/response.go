package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func ResponseCreateSecretResponse(ctx *fasthttp.RequestCtx, res SecretCreatedResponse){
	ctx.SetStatusCode(http.StatusOK)

	jres, err := json.Marshal(res)
	if err != nil {
		ResponseFatalError(ctx, err.Error())
		return
	}

	response(ctx, jres)
}

func ResponseApplicationError(ctx *fasthttp.RequestCtx, res ApplicationErrorResponse){
	jres, err := json.Marshal(res)
	if err != nil {
		ResponseFatalError(ctx, err.Error())
		return
	}

	response(ctx, jres)
}

func ResponseSecretReadResponse(ctx *fasthttp.RequestCtx, response []byte){
	ctx.SetStatusCode(http.StatusOK)
	fmt.Fprintf(ctx, "%s", response)
}

func ResponseFatalError(ctx *fasthttp.RequestCtx, err string){
	fmt.Fprintf(ctx, "%s", err)
}

func response(ctx *fasthttp.RequestCtx, data []byte) {
	if _, err := fmt.Fprintf(ctx, "%s", data); err != nil {
		log.Println(err)
	}
}
