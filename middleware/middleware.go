package middleware

import (
	"enterprise-hello-world/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const heartbeatPath = "/ping"

func InstallMiddleware(chiRouter chi.Router, conf *config.Configuration) {
	chiRouter.Use(
		middleware.Logger,
		middleware.Heartbeat(heartbeatPath),
		middleware.BasicAuth(
			"User authentication",
			map[string]string{
				conf.Username: conf.Password,
			},
		),
	)
}
