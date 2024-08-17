package main

import (
	"log"
	mongosh "user-service/internal/user/pkg/Mongosh"
	redis "user-service/internal/user/pkg/Redis"
	"user-service/internal/user/pkg/load"
	userservice "user-service/internal/user/pkg/user-service"
	"user-service/internal/user/repository"
	"user-service/internal/user/service"
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

	rdb, err := redis.InitRedis(*conf)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepo(mongoConn.Coll, rdb)
	service := service.NewService(userRepo)
	runService := userservice.NewRunSerivce(*service)

	log.Printf("User Service Running on :%d port", conf.UserServicePort)
	if err := runService.RUN(*conf); err != nil {
		log.Fatal(err)
	}
}
