package model

import (
	"regexp"
	"time"

	"github.com/google/uuid"
)

const signInAdminFailureCountMax = 4
const accountAdminLockedDulation = 30 * time.Minute

type User struct {
	id                 uuid.UUID `json:"id"`
	firstName          string    `json:"first_name"`
	lastName           string    `json:"last_name"`
	email              string    `json:"email"`
	passwordHash       PasswordHash
	signInFailedCount  int8
	accountLockedUntil *time.Time
}

type NewUserOption func(*User)

func NewUser(opts ...NewUserOption) (*User, error) {
	u := &User{}
	// member作成時にデフォルトでuuidを設定。上書き可能。
	u.id = uuid.Must(uuid.NewRandom())
	for _, opt := range opts {
		opt(u)
	}
	if err := u.validate(); err != nil {
		return nil, NewValidationError(err.Error())
	}
	return u, nil
}

func NewUserID(id uuid.UUID) NewUserOption {
	return func(u *User) {
		u.id = id
	}
}

func NewUserFirstName(firstName string) NewUserOption {
	return func(u *User) {
		u.firstName = firstName
	}
}

func NewUserLastName(lastName string) NewUserOption {
	return func(u *User) {
		u.lastName = lastName
	}
}

func NewUserEmail(email string) NewUserOption {
	return func(u *User) {
		u.email = email
	}
}

// domain rogic
func (u *User) IsAccountLocked() bool {
	return u.accountLockedUntil != nil && u.accountLockedUntil.After(time.Now())
}

// domain rogic
func (u *User) handleSignInFailure() {
	u.signInFailedCount++
	if u.signInFailedCount >= signInAdminFailureCountMax {
		u.signInFailedCount = 0
		accountLockedUntil := time.Now().Add(accountAdminLockedDulation)
		u.accountLockedUntil = &accountLockedUntil
	}
}

// domain rogic
func (u *User) SignIn(password string) error {
	err := u.passwordHash.ComparePassword(password)
	if err != nil {
		u.handleSignInFailure()
	} else {
		u.signInFailedCount = 0
	}
	return err
}

func (u *User) GetID() uuid.UUID {
	return u.id
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetSignInFailedCount() int8 {
	return u.signInFailedCount
}

func (u *User) GetAccountLockedUntil() *time.Time {
	return u.accountLockedUntil
}

func (u *User) isIDValid() *ValidationError {
	if u.id == uuid.Nil {
		return NewValidationError("empty UUID in member ID is not allowed")
	}
	return nil
}

func (u *User) isFirstNameValid() *ValidationError {
	if u.firstName == "" {
		return NewValidationError("empty first name is not allowed")
	}
	return nil
}

func (u *User) isLastNameValid() *ValidationError {
	if u.lastName == "" {
		return NewValidationError("empty last name is not allowed")
	}
	return nil
}

func (u *User) isEmailValid() *ValidationError {
	if u.email == "" {
		return NewValidationError("empty email is not allowed")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.email) {
		return NewValidationError("invalid email format")
	}
	return nil
}

func (u *User) isSignInFailedCountValid() *ValidationError {
	if u.signInFailedCount >= 0 && u.signInFailedCount <= signInAdminFailureCountMax {
		return nil
	}
	return NewValidationError("無効なログイン失敗回数です。")
}

func (u *User) isAccountLockedUntilValid() *ValidationError {
	if u.accountLockedUntil == nil {
		return nil
	}
	if u.accountLockedUntil.IsZero() {
		return NewValidationError("無効なアカウントロック期間です。")
	}
	return nil
}

func (u *User) validate() *ValidationErrors {
	var errors []*ValidationError
	if err := u.isIDValid(); err != nil {
		errors = append(errors, err)
	}
	if err := u.isFirstNameValid(); err != nil {
		errors = append(errors, err)
	}
	if err := u.isLastNameValid(); err != nil {
		errors = append(errors, err)
	}
	if err := u.isEmailValid(); err != nil {
		errors = append(errors, err)
	}
	if err := u.isSignInFailedCountValid(); err != nil {
		errors = append(errors, err)
	}
	if err := u.isAccountLockedUntilValid(); err != nil {
		errors = append(errors, err)
	}
	return validationErrorSliceToValidationErrors(errors)
}
