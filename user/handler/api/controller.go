package api

import (
	"miniproject/user/domain"
)

type userHandler struct {
	service domain.Service
}

func NewUserHandler(service domain.Service) userHandler {
	return userHandler{
		service: service,
	}
}
