package posts_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *PostRepository) GetPost(
	ctx context.Context,
	id int,
) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id,
		   version,
		   author_id,
		   category_id,
		   slug,
		   title,
		   content,
		   cover_image_ref,
		   status,
		   views_count,
		   created_at,
		   updated_at
	FROM posts
	WHERE id=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
	)

	var postModel PostModel

	err := row.Scan(
		&postModel.ID,
		&postModel.Version,
		&postModel.AuthorID,
		&postModel.CategoryID,
		&postModel.Slug,
		&postModel.Title,
		&postModel.Content,
		&postModel.CoverImageRef,
		&postModel.Status,
		&postModel.ViewsCount,
		&postModel.CreatedAt,
		&postModel.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Post{}, fmt.Errorf(
				"select post with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.Post{}, fmt.Errorf("scan error: %w", err)
	}

	postDomain := postDomainFromModel(postModel)

	return postDomain, nil
}
