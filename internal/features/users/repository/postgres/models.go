package users_postgres_repository

import (
	"time"

	"github.com/classytoast/go-blog/internal/core/domain"
)

type UserModel struct {
	ID           int64      `db:"id"`
	Version      int        `db:"version"`
	Username     string     `db:"username"`
	Email        string     `db:"email"`
	PasswordHash string     `db:"password_hash"`
	AvatarRef    *string    `db:"avatar_ref"`
	Bio          *string    `db:"bio"`
	Role         string     `db:"role"`
	IsActive     bool       `db:"is_active"`
	LastLoginAt  *time.Time `db:"last_login_at"`
	CreatedAt    time.Time  `db:"created_at"`
}

type UserToken struct {
	TokenHash string    `db:"token_hash"`
	UserID    int64     `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.User{
			ID:        user.ID,
			Version:   user.Version,
			Username:  user.Username,
			Email:     user.Email,
			AvatarRef: user.AvatarRef,
			Bio:       user.Bio,
			Role:      user.Role,
		}
	}

	return userDomains
}
