package bootstrap

import (
	"enterprise-hello-world/config"
	"enterprise-hello-world/middleware"
	"enterprise-hello-world/routers"
	"enterprise-hello-world/service"
	"github.com/go-chi/chi/v5"
	"log"
)

var l = log.Default()

func Setup(chiRouter chi.Router, conf *config.Configuration) {
	l.Printf(
		"Listening on port %v; "+
			"Message receiver is %v; "+
			"Username is %v; "+
			"Password is %v;",
		conf.Port,
		conf.MessageReceiver,
		conf.Username,
		conf.Password,
	)

	middleware.InstallMiddleware(chiRouter, conf)

	routers.NewHelloWorldRouter(
		chiRouter,
		service.NewHelloMessageService(conf),
	).Route()
}
