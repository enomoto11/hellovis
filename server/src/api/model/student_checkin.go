package model

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type StudentCheckin struct {
	id          uuid.UUID `json:"id"`
	studentID   uuid.UUID `json:"student_id"`
	manavisCode string    `json:"manavis_code"`
	checkedInAt time.Time `json:"checked_in_at"`
}

type NewStudentCheckinOption func(*StudentCheckin)

func NewStudentCheckin(opts ...NewStudentCheckinOption) (*StudentCheckin, error) {
	s := &StudentCheckin{}
	// member作成時にデフォルトでuuidを設定。上書き可能。
	s.id = uuid.Must(uuid.NewRandom())
	for _, opt := range opts {
		opt(s)
	}
	if err := s.validate(); err != nil {
		return nil, NewValidationError(err.Error())
	}
	return s, nil
}

func NewStudentCheckinID(id uuid.UUID) NewStudentCheckinOption {
	return func(s *StudentCheckin) {
		s.id = id
	}
}

func NewStudentCheckinStudentID(studentID uuid.UUID) NewStudentCheckinOption {
	return func(s *StudentCheckin) {
		s.studentID = studentID
	}
}

func NewStudentCheckinManavisCode(manavisCode string) NewStudentCheckinOption {
	return func(s *StudentCheckin) {
		s.manavisCode = manavisCode
	}
}

func NewStudentCheckinCheckedInAt(checkedInAt time.Time) NewStudentCheckinOption {
	return func(s *StudentCheckin) {
		s.checkedInAt = checkedInAt
	}
}

func (s *StudentCheckin) GetID() uuid.UUID {
	return s.id
}

func (s *StudentCheckin) GetStudentID() uuid.UUID {
	return s.studentID
}

func (s *StudentCheckin) GetManavisCode() string {
	return s.manavisCode
}

func (s *StudentCheckin) GetCheckedInAt() time.Time {
	return s.checkedInAt
}

func (sco *StudentCheckin) isIDValid() *ValidationError {
	if sco.id == uuid.Nil {
		return NewValidationError("empty UUID in member ID is not allowed")
	}
	return nil
}

func (sco *StudentCheckin) isStudentIDValid() *ValidationError {
	if sco.studentID == uuid.Nil {
		return NewValidationError("empty UUID in student ID is not allowed")
	}
	return nil
}

func (sco *StudentCheckin) isManavisCodeValid() *ValidationError {
	if sco.manavisCode == "" {
		return NewValidationError("empty manavis code is not allowed")
	}
	num, err := strconv.Atoi(sco.manavisCode)
	if err != nil {
		return NewValidationError("manavis code must be a number")
	}
	if num <= 0 {
		return NewValidationError("manavis code must be positive")
	}
	return nil
}

func (sco *StudentCheckin) isCheckedInAtValid() *ValidationError {
	if sco.checkedInAt.IsZero() {
		return NewValidationError("empty checked in at is not allowed")
	}
	return nil
}

func (s *StudentCheckin) validate() *ValidationErrors {
	var errors []*ValidationError
	if err := s.isIDValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isStudentIDValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isManavisCodeValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isCheckedInAtValid(); err != nil {
		errors = append(errors, err)
	}

	return validationErrorSliceToValidationErrors(errors)
}
