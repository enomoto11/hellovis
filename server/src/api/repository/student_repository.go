package repository

import (
	"context"
	"hellovis/api/model"
	"hellovis/api/utils"
	"hellovis/ent"
	"hellovis/ent/student"
	"hellovis/ent/studentcheckin"
	"hellovis/ent/studentcheckout"
	"time"

	"github.com/google/uuid"
)

type StudentRepository interface {
	Create(ctx *context.Context, s *model.Student) (*model.Student, error)
	FindByID(ctx *context.Context, id uuid.UUID) (*model.Student, error)
	FindByManavisCode(ctx *context.Context, manavisCode string) (*model.Student, error)
	FindAllByGradeAndIsInHigh(ctx *context.Context, grade int, isInHigh bool, pageable model.Pageable) (model.StudentPage, error)
	FindAllWhoHadCheckedInWithDayOffest(ctx *context.Context, dayOffset int) ([]*model.Student, error)
	FindAllWhoHaveCheckedInAndNotCheckedOut(ctx *context.Context) ([]*model.Student, error)
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

func (sr *studentRepository) FindAllByGradeAndIsInHigh(ctx *context.Context, grade int, isInHigh bool, pageable model.Pageable) (model.StudentPage, error) {
	builder := sr.client.Student.Query().
		Where(student.GradeEQ(grade), student.IsHighSchoolEQ(isInHigh))

		// 合計数取得
	total, totalErr := builder.Count(*ctx)
	if totalErr != nil {
		return nil, totalErr
	}

	entities, err := builder.
		Order(ent.Desc(student.FieldUpdatedAt)).
		Offset(pageable.GetOffset()).
		Limit(pageable.GetPageSize()).
		All(*ctx)
	if err != nil {
		return nil, err
	}

	models, err := utils.MapSliceWithError(entities, newStudentFromEntity)
	if err != nil {
		return nil, err
	}

	return model.NewPage(
		models,
		total,
		pageable.GetPageNumber(),
		pageable.GetPageSize(),
	), nil
}

// FindAllWhoHasCheckedInToday returns all students who have checked in today
// あくまでその日にチェックインした生徒を返すだけで、その時刻の取得は別の責務に任せる
// dayOffet 日前の来校済み生徒を返す
func (sr *studentRepository) FindAllWhoHadCheckedInWithDayOffest(ctx *context.Context, dayOffset int) ([]*model.Student, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location()).AddDate(0, 0, (-1)*dayOffset)
	endOfDay := startOfDay.AddDate(0, 0, 1)

	entities, err := sr.client.Student.Query().
		Where(student.HasCheckinsWith(studentcheckin.CheckinAtGTE(startOfDay), studentcheckin.CheckinAtLT(endOfDay))).
		All(*ctx)
	if err != nil {
		return nil, err
	}

	return utils.MapSliceWithError(entities, newStudentFromEntity)
}

func (sr *studentRepository) FindAllWhoHaveCheckedInAndNotCheckedOut(ctx *context.Context) ([]*model.Student, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)

	entities, err := sr.client.Student.Query().
		Where(student.HasCheckinsWith(studentcheckin.CheckinAtGTE(startOfDay), studentcheckin.CheckinAtLT(endOfDay))).
		Where(student.HasCheckoutsWith(studentcheckout.CheckoutAtLT(startOfDay))).
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
