package posts_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *PostRepository) PatchPost(
	ctx context.Context,
	id int,
	post domain.Post,
) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE posts
	SET 
		title=$1,
		content=$2,
		cover_image_ref=$3,
		status=$4,
		version=version + 1
	WHERE id=$5 AND version=$6

	RETURNING
		id,
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
		updated_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		post.Title,
		post.Content,
		post.CoverImageRef,
		post.Status,
		id,
		post.Version,
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
				"post with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Post{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	postDomain := postDomainFromModel(postModel)

	return postDomain, nil
}
