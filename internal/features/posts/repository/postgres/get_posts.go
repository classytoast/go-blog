package posts_postgres_repository

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (r *PostRepository) GetPosts(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Post, error) {
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
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	args := []any{limit, offset}

	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_id = $3")
		args = append(args, userID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select posts: %w", err)
	}
	defer rows.Close()

	var postsModels []PostModel

	for rows.Next() {
		var postModel PostModel

		err := rows.Scan(
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
			return nil, fmt.Errorf("scan posts: %w", err)
		}

		postsModels = append(postsModels, postModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	postsDomains := postDomainsFromModels(postsModels)

	return postsDomains, nil
}
