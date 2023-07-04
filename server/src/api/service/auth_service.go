package service

import (
	"context"
	"errors"
	"hellovis/api/repository"
	"net/http"
)

type AuthService interface {
	SignIn(ctx context.Context, params *SignInParams) (string, *InternalError)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
) AuthService {
	return &authService{
		userRepo,
	}
}

type SignInParams struct {
	Email    string
	Password string
}

// TOKEN を返す
func (s *authService) SignIn(ctx context.Context, params *SignInParams) (string, *InternalError) {
	user, err := s.userRepo.FindByEmail(&ctx, params.Email)
	if err != nil {
		return "", NewInternalError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return "", NewInternalError(http.StatusUnauthorized, errors.New("user not found"))
	}

	if user.IsAccountLocked() {
		return "", NewInternalError(http.StatusUnauthorized, errors.New("user account is locked"))
	}
	err = user.SignIn(params.Password)
	if err != nil {
		return "", NewInternalError(http.StatusUnauthorized, err)
	}
	if _, err := s.userRepo.Update(&ctx, user); err != nil {
		return "", NewInternalError(http.StatusInternalServerError, err)
	}

	return "tokenを返す", nil
}
