package repository

import (
	"context"
	"hellovis/api/model"
	"hellovis/ent"
)

type StudentCheckinRepository interface {
	Create(ctx *context.Context, s *model.StudentCheckin) (*model.StudentCheckin, error)
	Update(ctx *context.Context, s *model.StudentCheckin) (*model.StudentCheckin, error)
}

type studentCheckinRepository struct {
	client *ent.Client
}

func NewStudentCheckinRepository(client *ent.Client) StudentCheckinRepository {
	return &studentCheckinRepository{client}
}

func (scr *studentCheckinRepository) Create(ctx *context.Context, s *model.StudentCheckin) (*model.StudentCheckin, error) {
	entity, err := scr.client.StudentCheckin.Create().
		SetID(s.GetID()).
		SetStudentID(s.GetStudentID()).
		SetCheckinAt(s.GetCheckedInAt()).
		Save(*ctx)

	if err != nil {
		return nil, err
	}

	return newStudentCheckinFromEntity(entity)
}

// 使うことはあまり想定されていない
func (scr *studentCheckinRepository) Update(ctx *context.Context, s *model.StudentCheckin) (*model.StudentCheckin, error) {
	entity, err := scr.client.StudentCheckin.UpdateOneID(s.GetID()).
		SetStudentID(s.GetStudentID()).
		SetCheckinAt(s.GetCheckedInAt()).
		Save(*ctx)

	if err != nil {
		return nil, err
	}

	return newStudentCheckinFromEntity(entity)
}

func newStudentCheckinFromEntity(entity *ent.StudentCheckin) (*model.StudentCheckin, error) {
	return model.NewStudentCheckin(
		model.NewStudentCheckinID(entity.ID),
		model.NewStudentCheckinStudentID(entity.StudentID),
		model.NewStudentCheckinCheckedInAt(entity.CheckinAt),
	)
}
