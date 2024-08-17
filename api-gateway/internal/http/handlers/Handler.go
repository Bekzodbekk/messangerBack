package handlers

import "api-gateway/internal/service"

type HandlerST struct {
	service *service.ServiceRepositoryClient
}

func NewHandlerSt(service service.ServiceRepositoryClient) *HandlerST {
	return &HandlerST{
		service: &service,
	}
}
