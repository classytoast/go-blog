package posts_service

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (s *PostService) GetPost(
	ctx context.Context,
	id int,
) (domain.Post, error) {
	post, err := s.postRepository.GetPost(ctx, id)
	if err != nil {
		return domain.Post{}, fmt.Errorf(
			"get post from repository: %w", err,
		)
	}

	return post, nil
}
