package domain

import (
	"fmt"
	"strings"

	core_errors "github.com/classytoast/go-blog/internal/core/errors"
)

type User struct {
	ID        int64
	Version   int
	Username  string
	Email     string
	AvatarRef *string
	Bio       *string
	Role      string
}

type NewUserUninitialized struct {
	Username  string
	Email     string
	Password  string
	AvatarRef *string
	Bio       *string
}

func (u *User) Validate() error {
	if err := ValidateUsername(u.Username); err != nil {
		return err
	}

	if u.Bio != nil {
		if err := ValidateBio(*u.Bio); err != nil {
			return err
		}
	}

	if err := ValidateEmail(u.Email); err != nil {
		return err
	}

	return nil
}

func (u *NewUserUninitialized) Validate() error {
	if err := ValidateUsername(u.Username); err != nil {
		return err
	}

	if u.Bio != nil {
		if err := ValidateBio(*u.Bio); err != nil {
			return err
		}
	}

	if err := ValidateEmail(u.Email); err != nil {
		return err
	}

	return nil
}

func ValidateUsername(username string) error {
	usernameLength := len([]rune(username))
	if usernameLength < 3 || usernameLength > 100 {
		return fmt.Errorf(
			"invalid `Username` len: %d: %w",
			usernameLength,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func ValidateBio(bio string) error {
	bioLength := len([]rune(bio))
	if bioLength > 500 {
		return fmt.Errorf(
			"invalid `Bio` len: %s: %w",
			bioLength,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf(
			"invalid `Email`, undefined '@' symbol in %s: %w",
			email,
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

type UserPatch struct {
	Username  Nullable[string]
	Email     Nullable[string]
	AvatarRef Nullable[string]
	Bio       Nullable[string]
	Role      Nullable[string]
}

func NewUserPatch(
	Username Nullable[string],
	Email Nullable[string],
	AvatarRef Nullable[string],
	Bio Nullable[string],
	Role Nullable[string],
) UserPatch {
	return UserPatch{
		Username:  Username,
		Email:     Email,
		AvatarRef: AvatarRef,
		Bio:       Bio,
		Role:      Role,
	}
}

func (p *UserPatch) Validate() error {
	if p.Username.Set && p.Username.Value == nil {
		return fmt.Errorf(
			"`Username` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.Email.Set && p.Email.Value == nil {
		return fmt.Errorf(
			"`Email` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.Role.Set && p.Role.Value == nil {
		return fmt.Errorf(
			"`Role` can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch %w", err)
	}

	tmp := *u

	if patch.Username.Set {
		tmp.Username = *patch.Username.Value
	}
	if patch.Email.Set {
		tmp.Email = *patch.Email.Value
	}
	if patch.AvatarRef.Set {
		tmp.AvatarRef = patch.AvatarRef.Value
	}
	if patch.Bio.Set {
		tmp.Bio = patch.Bio.Value
	}
	if patch.Role.Set {
		tmp.Role = *patch.Role.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
