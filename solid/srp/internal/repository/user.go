package repository

import "srp/internal/domain"

type UserRepository struct {
	users []domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]domain.User, 0),
	}
}

func (r *UserRepository) GetAll() []domain.User {
	return r.users
}

func (r *UserRepository) Add(user *domain.User) *domain.User {
	r.users = append(r.users, *user)
	return user
}
