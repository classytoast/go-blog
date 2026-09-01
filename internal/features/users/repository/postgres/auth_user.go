package users_postgres_repository

// func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
// 	user := &domain.User{}

// 	err := r.pool.QueryRow(
// 		ctx, `
// 		SELECT
// 			id,
// 			email,
// 			password_hash
// 		FROM users
// 		WHERE email = $1
// 	`, email).Scan(
// 		&user.ID,
// 		&user.Email,
// 		&user.PasswordHash,
// 	)

// 	if err == sql.ErrNoRows {
// 		return nil, core_errors.ErrInvalidAuthData
// 	} else if err != nil {
// 		return nil, core_errors.ErrConnectDB
// 	}

// 	return user, nil
// }

// func (r *UserRepository) SaveToken(ctx context.Context, tokenModel *domain.UserToken) error {
// 	_, err := r.pool.Exec(
// 		ctx, `
// 		INSERT INTO users_tokens (
// 			token_hash,
// 			user_id,
// 			expires_at
// 		)
// 		VALUES ($1, $2, $3)
// 	`,
// 		&tokenModel.TokenHash,
// 		&tokenModel.UserID,
// 		&tokenModel.ExpiresAt,
// 	)

// 	return err
// }
