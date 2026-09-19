package posts_transport_http

import (
	"context"
	"net/http"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_http_server "github.com/classytoast/go-blog/internal/core/transport/http/server"
)

type PostService interface {
	CreatePost(
		ctx context.Context,
		post domain.Post,
	) (domain.Post, error)

	CreateCategory(
		ctx context.Context,
		category domain.Category,
	) (domain.Category, error)

	GetPosts(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Post, error)

	GetPost(
		ctx context.Context,
		id int,
	) (domain.Post, error)

	DeletePost(
		ctx context.Context,
		id int,
	) error

	PatchPost(
		ctx context.Context,
		id int,
		patch domain.PostPatch,
	) (domain.Post, error)
}

type PostHTTPHandler struct {
	postService PostService
}

func NewPostHTTPHandler(
	postService PostService,
) *PostHTTPHandler {
	return &PostHTTPHandler{
		postService: postService,
	}
}

func (h *PostHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/categories",
			Handler: h.CreateCategory,
		},
		{
			Method:  http.MethodPost,
			Path:    "/posts",
			Handler: h.CreatePost,
		},
		{
			Method:  http.MethodGet,
			Path:    "/posts",
			Handler: h.GetPosts,
		},
		{
			Method:  http.MethodGet,
			Path:    "/posts/{id}",
			Handler: h.GetPost,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/posts/{id}",
			Handler: h.DeletePost,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/posts/{id}",
			Handler: h.PatchPost,
		},
	}
}
