package posts_postgres_repository

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (r *PostRepository) CreateCategory(
	ctx context.Context,
	category domain.Category,
) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO categories (name)
	VALUES ($1)
	RETURNING id, name;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		category.Name,
	)

	var categoryModel CategoryModel

	err := row.Scan(
		&categoryModel.ID,
		&categoryModel.Name,
	)
	if err != nil {
		return domain.Category{}, fmt.Errorf("scan error: %w", err)
	}

	categoryDomain := domain.NewCategory(
		categoryModel.ID,
		category.Name,
	)

	return categoryDomain, nil
}
