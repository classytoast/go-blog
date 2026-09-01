package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, username, email, avatar_ref, bio, role
		FROM users
		WHERE id=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
	)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Username,
		&userModel.Email,
		&userModel.AvatarRef,
		&userModel.Bio,
		&userModel.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"select user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.User{
		ID:        userModel.ID,
		Version:   userModel.Version,
		Username:  userModel.Username,
		Email:     userModel.Email,
		AvatarRef: userModel.AvatarRef,
		Bio:       userModel.Bio,
		Role:      userModel.Role,
	}

	return userDomain, nil
}
