package main

import (
	"enterprise-hello-world/internal/bootstrapper"
	"enterprise-hello-world/internal/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var (
	conf      = config.ParseConfiguration()
	chiRouter = chi.NewRouter()
	l         = log.Default()
)

var data = []byte("the thing")

func main() {
	bootstrapper.Setup(chiRouter, conf)

	if err := http.ListenAndServe(":"+conf.Port, chiRouter); err != nil {
		l.Fatal(err)
	}
}
