package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func ResponseCreateSecretResponse(ctx *fasthttp.RequestCtx, res SecretCreatedResponse){
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

func ResponseFatalError(ctx *fasthttp.RequestCtx, err string){
	fmt.Fprintf(ctx, "%s", err)
}

func response(ctx *fasthttp.RequestCtx, data []byte) {
	if _, err := fmt.Fprintf(ctx, "%s", data); err != nil {
		log.Println(err)
	}
}
