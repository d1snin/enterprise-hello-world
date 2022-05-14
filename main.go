package main

import (
	"enterprise-hello-world/bootstrap"
	"enterprise-hello-world/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var (
	conf      = config.ParseConfiguration()
	chiRouter = chi.NewRouter()
	l         = log.Default()
)

func main() {
	bootstrap.Setup(chiRouter, conf)

	if err := http.ListenAndServe(":"+conf.Port, chiRouter); err != nil {
		l.Fatal(err)
	}
}
