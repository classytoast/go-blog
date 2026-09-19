package posts_service

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (s *PostService) PatchPost(
	ctx context.Context,
	id int,
	patch domain.PostPatch,
) (domain.Post, error) {
	post, err := s.postRepository.GetPost(ctx, id)
	if err != nil {
		return domain.Post{}, fmt.Errorf("get post: %w", err)
	}

	if err := post.ApplyPatch(patch); err != nil {
		return domain.Post{}, fmt.Errorf("apply post patch: %w", err)
	}

	patchedPost, err := s.postRepository.PatchPost(ctx, id, post)
	if err != nil {
		return domain.Post{}, fmt.Errorf("patch post: %w", err)
	}

	return patchedPost, nil
}
