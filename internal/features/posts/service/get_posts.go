package posts_service

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
)

func (s *PostService) GetPosts(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Post, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	posts, err := s.postRepository.GetPosts(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf(
			"get posts from repository: %w", err,
		)
	}

	return posts, nil
}
