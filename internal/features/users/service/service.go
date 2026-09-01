package users_service

import (
	"context"

	"github.com/classytoast/go-blog/internal/core/domain"
)

type UserRepository interface {
	UserExists(
		ctx context.Context,
		email string,
	) (bool, error)
	CreateUser(
		ctx context.Context,
		user *domain.NewUserUninitialized,
		passwordHash string,
	) error
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
