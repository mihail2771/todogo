package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/mihail2771/todogo/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(
	ID int,
	Version int,
	FullName string,
	PhoneNumber *string,
) User {
	return User{
		ID:          ID,
		Version:     Version,
		FullName:    FullName,
		PhoneNumber: PhoneNumber,
	}
}

func NewUserUninitialised(FullName string, PhoneNumber *string) User {
	return NewUser(
		UninitialisedID,
		UninitialisedVersion,
		FullName,
		PhoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid full name length %d:%w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		re := regexp.MustCompile(`^\+[0-9]{9,14}$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid phone number format %s:%w",
				*u.PhoneNumber,
				core_errors.ErrInvalidArgument,
			)

		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(
	fullName Nullable[string],
	phoneNumber Nullable[string],
) UserPatch {
	return UserPatch{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("FullName can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user %w", err)
	}

	*u = tmp

	return nil
}
