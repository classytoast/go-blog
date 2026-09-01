package users_transport_http

import (
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_http_types "github.com/classytoast/go-blog/internal/core/transport/http/types"
)

type NewUserDTORequest struct {
	Username string  `json:"username" validate:"required,min=3,max=100"`
	Email    string  `json:"email" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Avatar   *string `json:"avatar,omitempty"`
	Bio      *string `json:"bio,omitempty" validate:"omitempty,max=500"`
}

type AuthUserDTORequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTOResponse struct {
	ID        int64   `json:"id"`
	Version   int     `json:"version"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	AvatarRef *string `json:"avatar"`
	Bio       *string `json:"bio"`
	Role      string  `json:"role"`
}

type PatchUserRequest struct {
	Username  core_http_types.Nullable[string] `json:"username"`
	Email     core_http_types.Nullable[string] `json:"email"`
	AvatarRef core_http_types.Nullable[string] `json:"avatar"`
	Bio       core_http_types.Nullable[string] `json:"bio"`
	Role      core_http_types.Nullable[string] `json:"role"`
}

func UserPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.Username.ToDomain(),
		request.Email.ToDomain(),
		request.AvatarRef.ToDomain(),
		request.Bio.ToDomain(),
		request.Role.ToDomain(),
	)
}

func (r *PatchUserRequest) Validate() error {
	if r.Username.Set {
		if r.Username.Value == nil {
			return fmt.Errorf("`Username` can't be NULL")
		}

		usernameLen := len([]rune(*r.Username.Value))
		if usernameLen < 3 || usernameLen > 100 {
			return fmt.Errorf("`Username` must be between 3 and 100 symbols")
		}
	}

	if r.Bio.Set {
		if r.Bio.Value != nil {
			bioLen := len([]rune(*r.Bio.Value))
			if bioLen > 500 {
				return fmt.Errorf("`Bio` must be less than 500 symbols")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

func UserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:        user.ID,
		Version:   user.Version,
		Username:  user.Username,
		Email:     user.Email,
		AvatarRef: user.AvatarRef,
		Bio:       user.Bio,
		Role:      user.Role,
	}
}
