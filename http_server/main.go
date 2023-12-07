package main

import (
	"github.com/greeflas/itea_lessons/handler"
	"github.com/greeflas/itea_lessons/middleware"
	"github.com/greeflas/itea_lessons/server"
	"log"
)

func main() {
	logger := log.Default()

	requestLoggerMiddleware := middleware.NewRequestLogger(logger)

	homeHandler := handler.NewHomeHandler(logger)
	commentHandler := handler.NewCommentHandler(logger)

	apiServer := server.NewAPIServer(logger)
	apiServer.AddRoute("/", requestLoggerMiddleware.Wrap(homeHandler))
	apiServer.AddRoute("/comment", requestLoggerMiddleware.Wrap(commentHandler))

	if err := apiServer.Start(); err != nil {
		logger.Fatal(err)
	}
}
