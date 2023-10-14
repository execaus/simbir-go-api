package main

import (
	"context"
	"github.com/execaus/exloggo"
	"log"
	"os"
	"os/signal"
	"simbir-go-api/cache"
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
// @BasePath  /api/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	var serverInstance server.Server

	runLogger()

	env := models.LoadEnv()
	config := configs.LoadConfig()

	connection, queries := repository.NewBusinessDatabase(env, config)

	repos := repository.NewRepository(queries, connection)
	reposCache, err := loadCache(repos)
	if err != nil {
		exloggo.Fatal(err.Error())
	}
	services := service.NewService(repos, env, reposCache)
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

func runLogger() {
	if err := exloggo.SetParameters(&exloggo.Parameters{
		Directory: "logs",
	}); err != nil {
		log.Fatal(err.Error())
	}
}

func runChannelStopServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}

func loadCache(r *repository.Repository) (*cache.Cache, error) {
	c := cache.NewCache()

	dictionary, err := r.CacheBuilder.CacheRoles()
	if err != nil {
		exloggo.Error(err.Error())
		return nil, err
	}
	c.Role.Load(dictionary)

	return c, nil
}
