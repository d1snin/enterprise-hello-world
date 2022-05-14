package routers

import (
	"encoding/json"
	"enterprise-hello-world/routers/constant"
	"enterprise-hello-world/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HelloWorldRouter struct {
	Router
	chiRouter    chi.Router
	helloService service.MessageService
}

func NewHelloWorldRouter(chiRouter chi.Router, helloService service.MessageService) *HelloWorldRouter {
	return &HelloWorldRouter{
		helloService: helloService,
		chiRouter:    chiRouter,
	}
}

func (HelloWorldRouter *HelloWorldRouter) Route() {
	HelloWorldRouter.chiRouter.Get(
		constant.HelloWorldPath,

		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json")

			jsonEncoder := json.NewEncoder(writer)

			if err := jsonEncoder.Encode(
				HelloWorldRouter.helloService.NewCommonMessage(),
			); err != nil {
				log.Println(err)
			}
		},
	)
}
