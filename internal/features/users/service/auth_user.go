package users_service

// func (s *UserService) AuthUser(user *dto.AuthUserRequest) (string, error) {
// 	userDB, err := s.repo.FindByEmail(user.Email)
// 	if err != nil {
// 		return "", err
// 	}

// 	if !CheckPasswordHash(user.Password, userDB.PasswordHash) {
// 		return "", core_errors.ErrInvalidAuthData
// 	}

// 	token, err := GenerateToken()
// 	if err != nil {
// 		return "", core_errors.ErrGenerateToken
// 	}

// 	tokenHash := HashToken(token)

// 	tokenDB := domain.UserToken{
// 		TokenHash: tokenHash,
// 		UserID:    userDB.ID,
// 		ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
// 	}

// 	err = s.repo.SaveToken(&tokenDB)
// 	if err != nil {
// 		return "", core_errors.ErrConnectDB
// 	}

// 	return token, nil
// }

// // CheckPasswordHash сверяет чистый пароль с сохраненным хешем
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func GenerateToken() (string, error) {
// 	b := make([]byte, 32)

// 	_, err := rand.Read(b)
// 	if err != nil {
// 		return "", err
// 	}

// 	return hex.EncodeToString(b), nil
// }

// func HashToken(token string) string {
// 	sum := sha256.Sum256([]byte(token))
// 	return hex.EncodeToString(sum[:])
// }
