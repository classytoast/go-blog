package statistics_postgres_repository

import (
	"time"

	"github.com/classytoast/go-blog/internal/core/domain"
)

type PostModel struct {
	ID            int64     `db:"id"`
	Version       int       `db:"vesrion"`
	AuthorID      int64     `db:"author_id"`
	CategoryID    int64     `db:"category_id"`
	Slug          string    `db:"slug"`
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	CoverImageRef *string   `db:"cover_image_ref"`
	Status        string    `db:"status"`
	ViewsCount    int64     `db:"views_count"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func postDomainsFromModels(postModels []PostModel) []domain.Post {
	domains := make([]domain.Post, len(postModels))

	for i, model := range postModels {
		domains[i] = postDomainFromModel(model)
	}

	return domains
}

func postDomainFromModel(postModel PostModel) domain.Post {
	return domain.NewPost(
		postModel.ID,
		postModel.Version,
		postModel.AuthorID,
		postModel.CategoryID,
		postModel.Slug,
		postModel.Title,
		postModel.Content,
		postModel.CoverImageRef,
		postModel.Status,
		postModel.ViewsCount,
		postModel.CreatedAt,
		postModel.UpdatedAt,
	)
}
