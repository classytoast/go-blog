package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/classytoast/go-blog/internal/core/domain"
)

func (r *StatisticsRepository) GetPosts(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
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
	`)

	args := []any{}
	conditions := []string{}

	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("author_id=$%d", len(args)+1))
		args = append(args, userID)
	}

	if from != nil {
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)+1))
		args = append(args, from)
	}

	if to != nil {
		conditions = append(conditions, fmt.Sprintf("created_at<$%d", len(args)+1))
		args = append(args, to)
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	queryBuilder.WriteString(" ORDER BY id ASC;")

	rows, err := r.pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("select posts: %w", err)
	}
	defer rows.Close()

	var postModels []PostModel

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

		postModels = append(postModels, postModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	postsDomains := postDomainsFromModels(postModels)

	return postsDomains, nil
}
