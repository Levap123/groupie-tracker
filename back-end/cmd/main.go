package main

import (
	"groupie-tracker/back-end/internal/client"
	"groupie-tracker/back-end/internal/handler"
	"groupie-tracker/back-end/internal/service"
	"groupie-tracker/back-end/pkg/logger"
	"groupie-tracker/back-end/pkg/server"
	"log"
	"net/http"
)

func main() {
	client := client.NewClient(new(http.Client))
	service := service.NewService(client)
	logger := logger.NewLogger()
	handler := handler.NewHandler(service, logger)
	server := new(server.Server)

	logger.Info.Println("server is listening on: http://localhost:8080")
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
