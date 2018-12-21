package main

import (
	"gateway/routers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/*path", routers.Api)
	router.POST("/*path", routers.Api)
	router.PUT("/*path", routers.Api)
	router.DELETE("/*path", routers.Api)
	router.OPTIONS("/*path", routers.Cros)
	log.Fatal(http.ListenAndServe(":80", router))

}
