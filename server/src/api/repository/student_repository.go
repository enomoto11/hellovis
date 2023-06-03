package repository

import (
	"context"
	"hellovis/api/model"
	"hellovis/api/utils"
	"hellovis/ent"
	"hellovis/ent/student"

	"github.com/google/uuid"
)

type StudentRepository interface {
	Create(ctx *context.Context, s *model.Student) (*model.Student, error)
	FindByID(ctx *context.Context, id uuid.UUID) (*model.Student, error)
	FindByManavisCode(ctx *context.Context, manavisCode string) (*model.Student, error)
	FindAllByGradeAndIsInHigh(ctx *context.Context, grade int, isInHigh bool) ([]*model.Student, error)
}

type studentRepository struct {
	client *ent.Client
}

func NewStudentRepository(client *ent.Client) StudentRepository {
	return &studentRepository{client}
}

func (sr *studentRepository) Create(ctx *context.Context, s *model.Student) (*model.Student, error) {
	entity, err := sr.client.Student.Create().
		SetID(s.GetID()).
		SetFirstName(s.GetFirstName()).
		SetLastName(s.GetLastName()).
		SetGrade(s.GetGrade()).
		SetManavisCode(s.GetManavisCode()).
		SetIsHighSchool(s.GetIsInHighSchool()).
		Save(*ctx)

	if err != nil {
		return nil, err
	}

	return newStudentFromEntity(entity)
}

func (sr *studentRepository) FindByID(ctx *context.Context, id uuid.UUID) (*model.Student, error) {
	entity, err := sr.client.Student.Query().
		Where(student.IDEQ(id)).
		Only(*ctx)

	if err != nil {
		return nil, err
	}

	return newStudentFromEntity(entity)
}

func (sr *studentRepository) FindByManavisCode(ctx *context.Context, manavisCode string) (*model.Student, error) {
	entity, err := sr.client.Student.Query().
		Where(student.ManavisCodeEQ(manavisCode)).
		Only(*ctx)
	if err != nil {
		return nil, err
	}

	return newStudentFromEntity(entity)
}

func (sr *studentRepository) FindAllByGradeAndIsInHigh(ctx *context.Context, grade int, isInHigh bool) ([]*model.Student, error) {
	entities, err := sr.client.Student.Query().
		Where(student.GradeEQ(grade), student.IsHighSchoolEQ(isInHigh)).
		All(*ctx)
	if err != nil {
		return nil, err
	}

	return utils.MapSliceWithError(entities, newStudentFromEntity)
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
