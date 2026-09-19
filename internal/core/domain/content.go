package domain

import "time"

type Comment struct {
	Id         int64     `db:"id" json:"id"`
	PostId     int64     `db:"post_id" json:"post_id"`
	ParentId   *int64    `db:"parent_id" json:"parent_id,omitempty"`
	UserId     int64     `db:"user_id" json:"user_id"`
	Content    string    `db:"content" json:"content"`
	LikesCount int64     `db:"likes_count" json:"likes_count"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
