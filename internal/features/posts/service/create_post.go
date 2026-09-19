package posts_service

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (s *PostService) CreatePost(
	ctx context.Context,
	post domain.Post,
) (domain.Post, error) {
	if err := post.Validate(); err != nil {
		return domain.Post{}, fmt.Errorf(
			"validate post domain: %w", err,
		)
	}

	post, err := s.postRepository.CreatePost(ctx, post)
	if err != nil {
		return domain.Post{}, fmt.Errorf(
			"create post: %w", err,
		)
	}

	return post, nil
}
