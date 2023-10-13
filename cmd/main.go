package main

import (
	"context"
	"github.com/execaus/exloggo"
	"os"
	"os/signal"
	"simbir-go-api/configs"
	_ "simbir-go-api/docs"
	"simbir-go-api/handler"
	"simbir-go-api/models"
	"simbir-go-api/repository"
	"simbir-go-api/server"
	"simbir-go-api/service"
	"syscall"
)

// @title           SimbirGoAPI
// @version         1.0.0
// @description     API for transportation rental service

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath  /api/
func main() {
	var serverInstance server.Server

	env := models.LoadEnv()
	config := configs.LoadConfig()

	database := repository.NewBusinessDatabase(env, config)

	repos := repository.NewRepository(database)
	services := service.NewService(repos, env)
	handlers := handler.NewHandler(services)

	go runServer(&serverInstance, handlers, config.Server)

	runChannelStopServer()

	serverInstance.Shutdown(context.Background())
}

func runServer(server *server.Server, handlers *handler.Handler, config *configs.ServerConfig) {
	ginEngine := handlers.InitRoutes()

	if err := server.Run(config.Port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			exloggo.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}
}

func runChannelStopServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}
