package app

import (
	"fmt"
	"github.com/valyala/fasthttp"
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

func TestRouter_GetShouldReturnCorrectFunc(t *testing.T) {
	router := NewRouter()

	handler := TestHandler{}
	router.LinkHandler("POST", "/test", handler)

	if router.Get("POST", "/test") == nil {
		t.Error("router should not return nil")
	}
}

func TestRouter_LinkShouldBindMethodAndPathWithHandler(t *testing.T) {
	router := NewRouter()

	handler := TestHandler{}
	router.LinkHandler("POST", "/test", handler)

	if router.route["POST"]["/test"] == nil {
		t.Error("router should bind method, path with handler")
	}
}

func TestRouter_LinkShouldBindMethodAndPathWithFunc(t *testing.T){
	router := NewRouter()

	f := func(ctx *fasthttp.RequestCtx){}
	router.LinkFunc("POST", "/test", f)
	if (router.route["POST"]["/test"]) == nil {
		t.Error("router should bind method, path with func")
	}
}

func TestDeepEqualFunc(t *testing.T){
	f := func(ctx *fasthttp.RequestCtx){}
	if fmt.Sprintf("%v", &f) != fmt.Sprintf("%v", &f) {
		t.Error("cannot use reflect.DeepEqual to compare between function")
	}
}