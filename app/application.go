package app

import (
	"fastvault/app/controllers/secret"
	"github.com/tspn/whenThis"
	"github.com/valyala/fasthttp"
	"log"
)

const (
	ENV_CONFIG_LOCATION = "CONFIG"
)

func Start() {
	InitialConfig()

	router := NewRouter()

	//Register Path
	router.LinkFunc("POST", "/secret", secret.Post)

	//start fasthttp application
	port := whenThis.
		IsEmptyString(Config.Port).
		UseThisString(":8080")
	log.Println("server start @ port", port)
	fasthttp.ListenAndServe(
		port,
		router.GetModule())
}
