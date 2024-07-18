package main

import (
	"fmt"
	"srp/internal/domain"
	"srp/internal/handler"
	"srp/internal/repository"
	"srp/internal/service"
)

func main() {
	fmt.Println("===== WELCOME =====")

	userRepo := repository.NewUserRepository()
	authService := service.NewAuthenticationService(*userRepo)
	authHandler := handler.NewAuthenticationHandler(*authService)

	userReq := domain.UserReq{
		Username: "nadiannis",
		Password: "password",
	}
	authHandler.Authenticate(userReq)

	authHandler.Register(userReq)
	authHandler.Authenticate(userReq)
}
