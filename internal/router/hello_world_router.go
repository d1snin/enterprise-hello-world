package router

import (
	"encoding/json"
	"enterprise-hello-world/internal/constants"
	"enterprise-hello-world/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// HelloWorldRouter is a Router implementation that routes the request to the main functionality.
type HelloWorldRouter struct {
	chiRouter    chi.Router
	helloService service.MessageService
}

// NewHelloWorldRouter creates a new HelloWorldRouter instance.
func NewHelloWorldRouter(chiRouter chi.Router, helloService service.MessageService) Router {
	return &HelloWorldRouter{
		helloService: helloService,
		chiRouter:    chiRouter,
	}
}

// Route routes the request to the main functionality,
// retrieves the CommonMessage instance and responds with it in JSON format.
func (router *HelloWorldRouter) Route() {
	router.chiRouter.Get(
		constants.HelloWorldPath,

		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json")

			jsonEncoder := json.NewEncoder(writer)

			if err := jsonEncoder.Encode(
				router.helloService.NewCommonMessage(),
			); err != nil {
				log.Println(err)
			}
		},
	)
}
