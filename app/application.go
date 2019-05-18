package app

import (
	"fastvault/app/configurations"
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
	secret.Initialise()

	router := NewRouter()

	//Register Path
	router.LinkFunc("POST", "/secret", secret.Post)

	//start fasthttp application
	port := whenThis.
		IsEmptyString(configurations.Current.Port).
		UseThisString(":8080")

	log.Println("server start @ port", port)
	err := fasthttp.ListenAndServe(
		port,
		router.GetModule())
	if err != nil {
		log.Fatal(err)
	}
}
