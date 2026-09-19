package posts_postgres_repository

import core_postgres_pool "github.com/classytoast/go-blog/internal/core/repository/postgres/pool"

type PostRepository struct {
	pool core_postgres_pool.Pool
}

func NewPostRepository(
	pool core_postgres_pool.Pool,
) *PostRepository {
	return &PostRepository{
		pool: pool,
	}
}
