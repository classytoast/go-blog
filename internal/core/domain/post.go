package domain

import (
	"fmt"
	"slices"
	"strings"
	"time"

	core_errors "github.com/classytoast/go-blog/internal/core/errors"
)

type Post struct {
	ID      int64
	Version int

	AuthorID   int64
	CategoryID int64

	Slug          string
	Title         string
	Content       string
	CoverImageRef *string
	Status        string
	ViewsCount    int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewPost(
	id int64,
	version int,
	authorID int64,
	categoryID int64,
	slug string,
	title string,
	content string,
	coverImageRef *string,
	status string,
	viewsCount int64,
	createdAt time.Time,
	updatedAt time.Time,
) Post {
	return Post{
		ID:            id,
		Version:       version,
		AuthorID:      authorID,
		CategoryID:    categoryID,
		Slug:          slug,
		Title:         title,
		Content:       content,
		CoverImageRef: coverImageRef,
		Status:        status,
		ViewsCount:    viewsCount,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func NewPostUninitialized(
	authorID int64,
	categoryID int64,
	title string,
	content string,
	coverImageRef *string,
) Post {
	createdTime := time.Now()

	return NewPost(
		int64(UninitializedID),
		UninitializedVersion,
		authorID,
		categoryID,
		"",
		title,
		content,
		coverImageRef,
		"draft",
		0,
		createdTime,
		createdTime,
	)
}

var ValidStatuses = []string{
	"draft",
	"published",
	"archived",
}

func (p *Post) Validate() error {
	titleLen := len([]rune(p.Title))
	if titleLen < 3 || titleLen > 200 {
		return fmt.Errorf(
			"invalid `title` len: %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}

	contentLen := len([]rune(p.Content))
	if contentLen < 1 || contentLen > 15000 {
		return fmt.Errorf(
			"invalid `content` len: %d: %w",
			contentLen,
			core_errors.ErrInvalidArgument,
		)
	}

	if !slices.Contains(ValidStatuses, p.Status) {
		return fmt.Errorf("status` must be one of: %s", strings.Join(ValidStatuses, ", "))
	}

	return nil
}

type PostPatch struct {
	Title         Nullable[string]
	Content       Nullable[string]
	CoverImageRef Nullable[string]
	Status        Nullable[string]
}

func NewPostPatch(
	title Nullable[string],
	content Nullable[string],
	coverImageRef Nullable[string],
	status Nullable[string],
) PostPatch {
	return PostPatch{
		Title:         title,
		Content:       content,
		CoverImageRef: coverImageRef,
		Status:        status,
	}
}

func (p *PostPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf(
			"`Title` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Content.Set && p.Content.Value == nil {
		return fmt.Errorf(
			"`Content` can't be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if p.Status.Set && p.Status.Value == nil {
		return fmt.Errorf(
			"`Status` can't be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (p *Post) ApplyPatch(patch PostPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate post patch: %w", err)
	}

	tmp := *p

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Content.Set {
		tmp.Content = *patch.Content.Value
	}

	if patch.CoverImageRef.Set {
		tmp.CoverImageRef = patch.CoverImageRef.Value
	}

	if patch.Status.Set {
		tmp.Status = *patch.Status.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched post: %w", err)
	}

	*p = tmp

	return nil
}
