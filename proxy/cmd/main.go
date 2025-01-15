package main

import (
	"proxy/internal/controller"
	"proxy/internal/repo"
	"proxy/internal/service"
)

func main() {
	repo := repo.NewCash()
	service := service.New(repo)
	handler := controller.NewHandler(&service)
	r := controller.NewRouter(handler)
	server := controller.NewServer(":8080", r)
	go server.Serve()
	server.Shutdown()
}
