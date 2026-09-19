package posts_transport_http

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_http_types "github.com/classytoast/go-blog/internal/core/transport/http/types"
)

type CreatePostRequest struct {
	AuthorID      int64   `json:"author_id" validate:"required"`
	CategoryID    int64   `json:"category_id" validate:"required"`
	Title         string  `json:"title" validate:"required,min=3,max=200"`
	Content       string  `json:"content" validate:"required,max=15000"`
	CoverImageRef *string `json:"cover_image_ref,omitempty"`
}

type PostDTOResponse struct {
	ID            int64     `json:"id"`
	Version       int       `json:"version"`
	AuthorID      int64     `json:"author_id"`
	CategoryID    int64     `json:"category_id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CoverImageRef *string   `json:"cover_image_ref,omitempty"`
	Status        string    `json:"status"`
	ViewsCount    int64     `json:"views_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreatePostResponse PostDTOResponse

func postDTOFromDomain(post domain.Post) PostDTOResponse {
	return PostDTOResponse{
		ID:            post.ID,
		Version:       post.Version,
		AuthorID:      post.AuthorID,
		CategoryID:    post.CategoryID,
		Slug:          post.Slug,
		Title:         post.Title,
		Content:       post.Content,
		CoverImageRef: post.CoverImageRef,
		Status:        post.Status,
		ViewsCount:    post.ViewsCount,
		CreatedAt:     post.CreatedAt,
		UpdatedAt:     post.UpdatedAt,
	}
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1"`
}

type CreateCategoryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetPostsResponse []PostDTOResponse

func postsDTOsFromDomains(posts []domain.Post) []PostDTOResponse {
	dtos := make([]PostDTOResponse, len(posts))

	for i, post := range posts {
		dtos[i] = postDTOFromDomain(post)
	}

	return dtos
}

type PatchPostRequest struct {
	Title         core_http_types.Nullable[string] `json:"title"`
	Content       core_http_types.Nullable[string] `json:"content"`
	CoverImageRef core_http_types.Nullable[string] `json:"cover_image_ref"`
	Status        core_http_types.Nullable[string] `json:"status"`
}

func (r *PatchPostRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` can't be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 3 || titleLen > 200 {
			return fmt.Errorf("`Title` must be between 3 and 200 symbols")
		}
	}

	if r.Content.Set {
		if r.Content.Value == nil {
			return fmt.Errorf("`Content` can't be NULL")
		}

		contentLen := len([]rune(*r.Content.Value))
		if contentLen < 1 || contentLen > 15000 {
			return fmt.Errorf("`Content` must be between 1 and 15000 symbols")
		}
	}

	if r.Status.Set {
		if r.Status.Value == nil {
			return fmt.Errorf("`Status` can't be NULL")
		}

		if !slices.Contains(domain.ValidStatuses, *r.Status.Value) {
			return fmt.Errorf("`Status` must be one of: %s", strings.Join(domain.ValidStatuses, ", "))
		}
	}

	return nil
}

func postPatchFromRequest(request PatchPostRequest) domain.PostPatch {
	return domain.NewPostPatch(
		request.Title.ToDomain(),
		request.Content.ToDomain(),
		request.CoverImageRef.ToDomain(),
		request.Status.ToDomain(),
	)
}
