package middleware

import (
	"enterprise-hello-world/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const _heartbeatPath = "/ping"

// InstallMiddleware installs the middleware to the chi.Router instance.
func InstallMiddleware(chiRouter chi.Router, conf *config.Configuration) {
	chiRouter.Use(
		middleware.Logger,
		middleware.Heartbeat(_heartbeatPath),
		middleware.BasicAuth(
			"User authentication",
			map[string]string{
				conf.Username: conf.Password,
			},
		),
	)
}
