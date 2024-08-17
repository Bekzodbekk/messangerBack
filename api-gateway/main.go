package main

import (
	"api-gateway/internal/http"
	"api-gateway/internal/pkg/load"
	messageservice "api-gateway/internal/pkg/message-service"
	userservice "api-gateway/internal/pkg/user-service"
	"api-gateway/internal/service"
	"fmt"
	"log"
)

func main() {
	conf, err := load.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	userDial, err := userservice.DialWithUserService(*conf)
	if err != nil {
		log.Fatal(err)
	}

	messageDial, err := messageservice.DialWithMessageService(*conf)
	if err != nil {
		log.Fatal(err)
	}

	repoClient := service.NewServiceRepositoryClinet(userDial, messageDial)
	r := http.NewGin(*repoClient)

	target := fmt.Sprintf("%s:%d", conf.ApiGatewayHost, conf.ApiGatewayPort)
	if err := r.Run(target); err != nil {
		log.Fatal(err)
	}
}
