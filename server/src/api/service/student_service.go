package service

import (
	"context"
	"hellovis/api/model"
	"hellovis/api/repository"

	"github.com/google/uuid"
)

type StudentService interface {
	Create(ctx context.Context, params *CreateStudentParams) (*model.Student, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Student, error)
	FindByManavisCode(ctx context.Context, manavisCode string) (*model.Student, error)
	FindAllByGradeAndIsInHigh(ctx context.Context, params *FindAllByGradeAndIsInHighParams) ([]*model.Student, error)
	DeleteByIDAndManavisCode(ctx context.Context, params *DeleteByManavisCodeParams) error
}

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(
	studentRepo repository.StudentRepository,
) StudentService {
	return &studentService{
		studentRepo,
	}
}

type CreateStudentParams struct {
	FirstName    string
	LastName     string
	Grade        int
	IsHighSchool bool
	ManavisCode  string
}

type FindAllByGradeAndIsInHighParams struct {
	Grade        int
	IsHighSchool bool
}

type DeleteByManavisCodeParams struct {
	ID          uuid.UUID
	ManavisCode string
}

func (s *studentService) Create(ctx context.Context, params *CreateStudentParams) (*model.Student, error) {
	student, err := model.NewStudent(
		model.NewStudentID(uuid.New()),
		model.NewStudentFirstName(params.FirstName),
		model.NewStudentLastName(params.LastName),
		model.NewStudentGrade(params.Grade),
		model.NewStudentManavisCode(params.ManavisCode),
		model.NewStudentIsInHighSchool(params.IsHighSchool),
	)
	if err != nil {
		return nil, err
	}

	result, err := s.studentRepo.Create(&ctx, student)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *studentService) FindByID(ctx context.Context, id uuid.UUID) (*model.Student, error) {
	result, err := s.studentRepo.FindByID(&ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *studentService) FindByManavisCode(ctx context.Context, manavisCode string) (*model.Student, error) {
	result, err := s.studentRepo.FindByManavisCode(&ctx, manavisCode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *studentService) FindAllByGradeAndIsInHigh(ctx context.Context, params *FindAllByGradeAndIsInHighParams) ([]*model.Student, error) {
	results, err := s.studentRepo.FindAllByGradeAndIsInHigh(&ctx, params.Grade, params.IsHighSchool)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *studentService) DeleteByIDAndManavisCode(ctx context.Context, params *DeleteByManavisCodeParams) error {
	if err := s.studentRepo.DeleteByIDAndManavisCode(&ctx, params.ID, params.ManavisCode); err != nil {
		return err
	}
	return nil
}
