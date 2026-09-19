package posts_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *PostRepository) CreatePost(
	ctx context.Context,
	post domain.Post,
) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO posts (
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
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id, 
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
		post.AuthorID,
		post.CategoryID,
		post.Slug,
		post.Title,
		post.Content,
		post.CoverImageRef,
		post.Status,
		post.ViewsCount,
		post.CreatedAt,
		post.UpdatedAt,
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return domain.Post{}, fmt.Errorf(
					"%v: authorID=`%d`,categoryID=`%d`: %w",
					core_errors.ErrViolatesForeignKey,
					post.AuthorID,
					post.CategoryID,
					core_errors.ErrNotFound,
				)
			}
		}

		return domain.Post{}, fmt.Errorf("scan error: %w", err)
	}

	postDomain := postDomainFromModel(postModel)

	return postDomain, nil
}
