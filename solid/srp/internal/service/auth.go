package service

import (
	"errors"
	"srp/internal/domain"
	"srp/internal/repository"
)

type AuthenticationService struct {
	userRepo repository.UserRepository
}

func NewAuthenticationService(userRepo repository.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepo: userRepo,
	}
}

func (s *AuthenticationService) GetByUsername(username string) (*domain.User, error) {
	for _, user := range s.userRepo.GetAll() {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *AuthenticationService) Register(username, password string) (*domain.User, error) {
	if _, err := s.GetByUsername(username); err == nil {
		return nil, errors.New("user already exists")
	}

	user := domain.User{
		Username: username,
		Password: password,
	}
	return s.userRepo.Add(&user), nil
}

func (s *AuthenticationService) Authenticate(username, password string) (bool, error) {
	user, err := s.GetByUsername(username)
	if err != nil {
		return false, err
	}

	return user.Password == password, nil
}
