package repository

import (
	"context"
	"hellovis/api/model"
	"hellovis/ent"
	"hellovis/ent/user"
)

type UserRepository interface {
	FindByEmail(ctx *context.Context, email string) (*model.User, error)
	Update(ctx *context.Context, user *model.User) (*model.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client}
}

func (ur *userRepository) FindByEmail(ctx *context.Context, email string) (*model.User, error) {
	entity, err := ur.client.User.Query().
		Where(user.EmailEQ(email)).
		Only(*ctx)

	if err != nil {
		return nil, err
	}

	return newUserFromEntity(entity)
}

func (ur *userRepository) Update(ctx *context.Context, m *model.User) (*model.User, error) {
	update := ur.client.User.UpdateOneID(m.GetID()).
		SetEmail(m.GetEmail()).
		SetFirstName(m.GetFirstName()).
		SetLastName(m.GetLastName()).
		SetPasswordHash(m.GetPasswordHash().String()).
		SetSignInFailedCount(m.GetSignInFailedCount())
	if accountLoackedUntil := m.GetAccountLockedUntil(); accountLoackedUntil == nil {
		update.ClearAccountLockedUntil()
	} else {
		update.SetAccountLockedUntil(*accountLoackedUntil)
	}

	result, err := update.Save(*ctx)
	if err != nil {
		return nil, err
	}

	return newUserFromEntity(result)
}

func newUserFromEntity(entity *ent.User) (*model.User, error) {
	opts := []model.NewUserOption{
		model.NewUserID(entity.ID),
		model.NewUserFirstName(entity.FirstName),
		model.NewUserLastName(entity.LastName),
		model.NewUserEmail(entity.Email),
	}

	return model.NewUser(opts...)
}
