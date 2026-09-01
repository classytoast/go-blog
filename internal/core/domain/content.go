package domain

import "time"

type Post struct {
	Id            int64     `db:"id" json:"id"`
	Version       int       `db:"vesrion" json:"version"`
	AuthorId      int64     `db:"author_id" json:"author_id"`
	CategoryId    int64     `db:"category_id" json:"category_id"`
	Slug          string    `db:"slug" json:"slug"`
	Title         string    `db:"title" json:"title"`
	Content       string    `db:"content" json:"content"`
	CoverImageRef *string   `db:"cover_image_ref" json:"cover_image_ref,omitempty"`
	Status        string    `db:"status" json:"status"`
	ViewsCount    int64     `db:"views_count" json:"views_count"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type Category struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Comment struct {
	Id         int64     `db:"id" json:"id"`
	PostId     int64     `db:"post_id" json:"post_id"`
	ParentId   *int64    `db:"parent_id" json:"parent_id,omitempty"`
	UserId     int64     `db:"user_id" json:"user_id"`
	Content    string    `db:"content" json:"content"`
	LikesCount int64     `db:"likes_count" json:"likes_count"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
