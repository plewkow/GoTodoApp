package main

import (
	"draft-zadania-1/api"
	"draft-zadania-1/config"
	"draft-zadania-1/kafka"
	"draft-zadania-1/repo"
	spec "draft-zadania-1/spec"
	"log"
	//"draft-zadania-1/router"
	"draft-zadania-1/services"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config.InitDB()
	brokers := []string{"localhost:9093"}

	topicNames := []string{"todo-user", "todo-task"}

	if err := kafka.EnsureTopicExists(brokers, topicNames); err != nil {
		log.Fatalf("Kafka topic init error: %v", err)
	}
	kafkaProducer, err := kafka.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatalf("Kafka producer error: %v", err)
	}
	defer kafkaProducer.Close()
	userRepo := repo.NewUserRepository(config.DB)
	taskRepo := repo.NewTaskRepository(config.DB)
	userService := services.NewUserService(userRepo, kafkaProducer)
	taskService := services.NewTaskService(taskRepo, userRepo, kafkaProducer)
	userHandler := &api.UserHandler{Service: userService}
	taskHandler := &api.TaskHandler{Service: taskService}
	combined := api.NewCombinedHandler(userHandler, taskHandler)
	e := echo.New()
	apiGroup := e.Group("/api")
	e.Static("/api", "dist")
	e.Static("/swagger-ui", "swagger-ui")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	spec.RegisterHandlers(apiGroup, combined)
	e.Logger.Fatal(e.Start(":8080"))
}
