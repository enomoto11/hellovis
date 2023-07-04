package service

import (
	"context"
	"hellovis/api/model"
	"hellovis/api/repository"
	"net/http"

	"github.com/google/uuid"
)

type StudentService interface {
	Create(ctx context.Context, params *CreateStudentParams) (*model.Student, *InternalError)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Student, *InternalError)
	FindByManavisCode(ctx context.Context, manavisCode string) (*model.Student, *InternalError)
	FindAllByGradeAndIsInHigh(ctx context.Context, params *FindAllByGradeAndIsInHighParams) (model.StudentPage, *InternalError)
	FindAllWhoHadCheckedInWithDayOffest(ctx context.Context, params FindAllWhoHadCheckedInWithDayOffestParams) ([]*model.Student, *InternalError)
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
	PageNumber   int
}

type FindAllWhoHadCheckedInWithDayOffestParams struct {
	DayOffset int
}

func (s *studentService) Create(ctx context.Context, params *CreateStudentParams) (*model.Student, *InternalError) {
	student, err := model.NewStudent(
		model.NewStudentID(uuid.New()),
		model.NewStudentFirstName(params.FirstName),
		model.NewStudentLastName(params.LastName),
		model.NewStudentGrade(params.Grade),
		model.NewStudentManavisCode(params.ManavisCode),
		model.NewStudentIsInHighSchool(params.IsHighSchool),
	)
	if err != nil {
		return nil, NewInternalError(http.StatusBadRequest, err)
	}

	result, err := s.studentRepo.Create(&ctx, student)
	if err != nil {
		return nil, NewInternalError(http.StatusInternalServerError, err)
	}

	return result, nil
}

func (s *studentService) FindByID(ctx context.Context, id uuid.UUID) (*model.Student, *InternalError) {
	result, err := s.studentRepo.FindByID(&ctx, id)
	if err != nil {
		return nil, NewInternalError(http.StatusInternalServerError, err)
	}

	return result, nil
}

func (s *studentService) FindByManavisCode(ctx context.Context, manavisCode string) (*model.Student, *InternalError) {
	result, err := s.studentRepo.FindByManavisCode(&ctx, manavisCode)
	if err != nil {
		return nil, NewInternalError(http.StatusInternalServerError, err)
	}

	return result, nil
}

const (
	pageSize          = 10
	defaultPageNumber = 1
)

func (s *studentService) FindAllByGradeAndIsInHigh(ctx context.Context, params *FindAllByGradeAndIsInHighParams) (model.StudentPage, *InternalError) {
	if params.PageNumber == 0 {
		params.PageNumber = defaultPageNumber
	}
	pageable := model.NewPageable(params.PageNumber, pageSize)

	result, err := s.studentRepo.FindAllByGradeAndIsInHigh(&ctx, params.Grade, params.IsHighSchool, pageable)
	if err != nil {
		return nil, NewInternalError(http.StatusInternalServerError, err)
	}

	return result, nil
}

func (c *studentService) FindAllWhoHadCheckedInWithDayOffest(ctx context.Context, params FindAllWhoHadCheckedInWithDayOffestParams) ([]*model.Student, *InternalError) {
	results, err := c.studentRepo.FindAllWhoHadCheckedInWithDayOffest(&ctx, params.DayOffset)
	if err != nil {
		return nil, NewInternalError(http.StatusInternalServerError, err)
	}

	return results, nil
}
