package app

import (
	"github.com/valyala/fasthttp"
	"reflect"
	"testing"
)

type TestHandler struct{}

func (handler TestHandler) Handle(ctx *fasthttp.RequestCtx) {}

func TestRouter_GetShouldReturnNilIfTryGetUnlinkMethodOrPath(t *testing.T) {
	router := NewRouter()

	if router.Get("POST", "/test") != nil {
		t.Error("router should return nil")
	}
}

func TestRouter_GetShouldReturnCorrectHandler(t *testing.T) {
	router := NewRouter()

	handler := TestHandler{}
	router.Link("POST", "/test", handler)

	if !reflect.DeepEqual(handler, router.Get("POST", "/test")) {
		t.Error("router should return handler")
	}
}

func TestRouter_LinkShouldBindMethodAndPathWithHandler(t *testing.T) {
	router := NewRouter()

	handler := TestHandler{}
	router.Link("POST", "/test", handler)

	if !reflect.DeepEqual(handler, router.route["POST"]["/test"]) {
		t.Error("router cannot link handler")
	}
}
