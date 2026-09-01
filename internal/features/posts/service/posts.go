package posts_service

import (
	post_repo "github.com/classytoast/go-blog/internal/features/posts/repository/postgres"
)

type PostService struct {
	repo *post_repo.PostRepository
}

func NewPostService(repo *post_repo.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}
