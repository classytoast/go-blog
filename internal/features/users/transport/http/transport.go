package users_transport_http

import (
	"context"
	"net/http"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_http_server "github.com/classytoast/go-blog/internal/core/transport/http/server"
)

type UserService interface {
	RegisterUser(
		ctx context.Context,
		user *domain.NewUserUninitialized,
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
		user domain.UserPatch,
	) (domain.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.RegisterUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
