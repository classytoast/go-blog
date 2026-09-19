package posts_service

import (
	"context"
	"fmt"
)

func (s *PostService) DeletePost(
	ctx context.Context,
	id int,
) error {
	if err := s.postRepository.DeletePost(ctx, id); err != nil {
		return fmt.Errorf("delete post from repository: %w", err)
	}

	return nil
}
