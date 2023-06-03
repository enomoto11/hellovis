package repository

import (
	"context"
	"hellovis/api/model"
	"hellovis/ent"
)

type StudentRepository interface {
	Create(ctx context.Context, s *model.Student) (*model.Student, error)
}

type studentRepository struct {
	client *ent.Client
}

func NewStudentRepository(client *ent.Client) StudentRepository {
	return &studentRepository{client}
}

func (sr *studentRepository) Create(ctx context.Context, s *model.Student) (*model.Student, error) {
	entity, err := sr.client.Student.Create().
		SetID(s.GetID()).
		SetFirstName(s.GetFirstName()).
		SetLastName(s.GetLastName()).
		SetGrade(s.GetGrade()).
		SetManavisCode(s.GetManavisCode()).
		SetIsHighSchool(s.GetIsInHighSchool()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return newStudentFromEntity(entity)
}

func newStudentFromEntity(entity *ent.Student) (*model.Student, error) {
	opts := []model.NewStudentOption{
		model.NewStudentID(entity.ID),
		model.NewStudentFirstName(entity.FirstName),
		model.NewStudentLastName(entity.LastName),
		model.NewStudentGrade(entity.Grade),
		model.NewStudentManavisCode(entity.ManavisCode),
		model.NewStudentIsInHighSchool(entity.IsHighSchool),
	}

	return model.NewStudent(opts...)
}
