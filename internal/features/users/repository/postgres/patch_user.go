package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE users
		SET 
			username=$1,
			email=$2, 
			avatar_ref=$3, 
			bio=$4, 
			role=$5,
			version=version+1
		WHERE id=$6 AND version=$7
		RETURNING
			id,
			version,
			username,
			email,
			avatar_ref,
			bio,
			role;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.Email,
		user.AvatarRef,
		user.Bio,
		user.Role,
		id,
		user.Version,
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
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.User{}, fmt.Errorf(
			"scan error: %w", err,
		)
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
