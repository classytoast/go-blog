package posts_service

import (
	"context"

	"github.com/classytoast/go-blog/internal/core/domain"
)

type PostRepository interface {
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
		post domain.Post,
	) (domain.Post, error)
}

type PostService struct {
	postRepository PostRepository
}

func NewPostService(
	postRepository PostRepository,
) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}
