package model

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	var number, alphabet bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c) || unicode.IsLower(c):
			alphabet = true
		default:
		}
	}
	eightOrMore := utf8.RuneCountInString(password) >= 8
	sixteenUnder := utf8.RuneCountInString(password) < 16
	if eightOrMore && sixteenUnder && number && alphabet {
		return &Password{password}, nil
	}
	return nil, NewValidationError("不正なパスワードです")
}

func (p Password) GenerateHash() (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.value), 12)
	return NewPasswordHash(string(hash)), err
}

func (p Password) String() string {
	return p.value
}

type PasswordHash struct {
	value string
}

func NewPasswordHash(hash string) PasswordHash {
	return PasswordHash{hash}
}

func (p PasswordHash) String() string {
	return p.value
}

func (p PasswordHash) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(password))
	return err
}
