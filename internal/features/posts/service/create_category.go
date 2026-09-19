package posts_service

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (s *PostService) CreateCategory(
	ctx context.Context,
	category domain.Category,
) (domain.Category, error) {
	if err := category.Validate(); err != nil {
		return domain.Category{}, fmt.Errorf(
			"validate category domain: %w", err,
		)
	}

	category, err := s.postRepository.CreateCategory(ctx, category)
	if err != nil {
		return domain.Category{}, fmt.Errorf(
			"create category: %w", err,
		)
	}

	return category, nil
}
