package repository

import (
	"context"
	"hellovis/api/model"
	"hellovis/ent"
)

type StudentCheckoutRepository interface {
	Create(ctx *context.Context, s *model.StudentCheckout) (*model.StudentCheckout, error)
	Update(ctx *context.Context, s *model.StudentCheckout) (*model.StudentCheckout, error)
}

type studentCheckoutRepository struct {
	client *ent.Client
}

func NewStudentCheckoutRepository(client *ent.Client) StudentCheckoutRepository {
	return &studentCheckoutRepository{client}
}

func (scr *studentCheckoutRepository) Create(ctx *context.Context, s *model.StudentCheckout) (*model.StudentCheckout, error) {
	entity, err := scr.client.StudentCheckout.Create().
		SetID(s.GetID()).
		SetStudentID(s.GetStudentID()).
		SetCheckoutAt(s.GetCheckedoutAt()).
		Save(*ctx)

	if err != nil {
		return nil, err
	}

	return newStudentCheckoutFromEntity(entity)
}

// 使うことはあまり想定されていない
func (scr *studentCheckoutRepository) Update(ctx *context.Context, s *model.StudentCheckout) (*model.StudentCheckout, error) {
	entity, err := scr.client.StudentCheckout.UpdateOneID(s.GetID()).
		SetStudentID(s.GetStudentID()).
		SetCheckoutAt(s.GetCheckedoutAt()).
		Save(*ctx)

	if err != nil {
		return nil, err
	}

	return newStudentCheckoutFromEntity(entity)
}

func newStudentCheckoutFromEntity(entity *ent.StudentCheckout) (*model.StudentCheckout, error) {
	return model.NewStudentCheckout(
		model.NewStudentCheckoutID(entity.ID),
		model.NewStudentCheckoutStudentID(entity.StudentID),
		model.NewStudentCheckoutCheckedoutAt(entity.CheckoutAt),
	)
}
