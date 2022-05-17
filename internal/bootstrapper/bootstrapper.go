package bootstrapper

import (
	"enterprise-hello-world/internal/config"
	"enterprise-hello-world/internal/middleware"
	"enterprise-hello-world/internal/router"
	"enterprise-hello-world/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
)

var l = log.Default()

// Setup prints the provided configuration, configures the middleware and request routers.
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

	router.NewHelloWorldRouter(
		chiRouter,
		service.NewHelloMessageService(conf),
	).Route()
}
