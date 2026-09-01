package users_postgres_repository

import (
	"context"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
)

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *domain.NewUserUninitialized,
	passwordHash string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	_, err := r.pool.Exec(
		ctx, `
    INSERT INTO users (
        username,
		password_hash,
        email,
        bio,
        avatar_ref
    )
    VALUES ($1, $2, $3, $4, $5);
	`,
		user.Username,
		passwordHash,
		user.Email,
		user.Bio,
		user.AvatarRef,
	)

	return err
}

func (r *UserRepository) UserExists(ctx context.Context, email string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(
		ctx, `
		SELECT EXISTS(
            SELECT 1
            FROM users
            WHERE email = $1
        );
	`, email).Scan(&exists)

	if err != nil {
		return false, core_errors.ErrConnectDB
	}

	return exists, nil
}
