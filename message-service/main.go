package main

import (
	"log"
	mongosh "message-service/internal/user/pkg/Mongosh"
	"message-service/internal/user/pkg/load"
	messageservice "message-service/internal/user/pkg/message-service"
	"message-service/internal/user/repository"
	"message-service/internal/user/service"
)

func main() {

	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	mongoConn, err := mongosh.InitDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewMessageRepo(mongoConn.Coll)
	service := service.NewService(userRepo)
	runService := messageservice.NewRunSerivce(*service)

	log.Printf("Message Service Running on :%d port", conf.MessageServicePort)
	if err := runService.RUN(*conf); err != nil {
		log.Fatal(err)
	}
}
