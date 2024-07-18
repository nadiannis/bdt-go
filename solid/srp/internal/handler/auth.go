package handler

import (
	"fmt"
	"srp/internal/domain"
	"srp/internal/service"
)

type AuthenticationHandler struct {
	authService service.AuthenticationService
}

func NewAuthenticationHandler(authService service.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authService: authService,
	}
}

func (h *AuthenticationHandler) Register(r domain.UserReq) {
	user, err := h.authService.Register(r.Username, r.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("user '%s' registered successfully\n", user.Username)
}

func (h *AuthenticationHandler) Authenticate(r domain.UserReq) {
	isAuthenticated, err := h.authService.Authenticate(r.Username, r.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isAuthenticated {
		fmt.Println("password is incorrect")
		return
	}

	fmt.Println("authentication is successful")
}
