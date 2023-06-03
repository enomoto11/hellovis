package model

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type StudentCheckout struct {
	id           uuid.UUID `json:"id"`
	studentID    uuid.UUID `json:"student_id"`
	manavisCode  string    `json:"manavis_code"`
	checkedOutAt time.Time `json:"checked_in_at"`
}

type NewStudentCheckoutOption func(*StudentCheckout)

func NewStudentCheckout(opts ...NewStudentCheckoutOption) (*StudentCheckout, error) {
	s := &StudentCheckout{}
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

func NewStudentCheckoutID(id uuid.UUID) NewStudentCheckoutOption {
	return func(s *StudentCheckout) {
		s.id = id
	}
}

func NewStudentCheckoutStudentID(studentID uuid.UUID) NewStudentCheckoutOption {
	return func(s *StudentCheckout) {
		s.studentID = studentID
	}
}

func NewStudentCheckoutManavisCode(manavisCode string) NewStudentCheckoutOption {
	return func(s *StudentCheckout) {
		s.manavisCode = manavisCode
	}
}

func NewStudentCheckooutCheckedOutAt(checkedOutAt time.Time) NewStudentCheckoutOption {
	return func(s *StudentCheckout) {
		s.checkedOutAt = checkedOutAt
	}
}

func (s *StudentCheckout) GetID() uuid.UUID {
	return s.id
}

func (s *StudentCheckout) GetStudentID() uuid.UUID {
	return s.studentID
}

func (s *StudentCheckout) GetManavisCode() string {
	return s.manavisCode
}

func (s *StudentCheckout) GetCheckedOutAt() time.Time {
	return s.checkedOutAt
}

func (sco *StudentCheckout) isIDValid() *ValidationError {
	if sco.id == uuid.Nil {
		return NewValidationError("empty UUID in member ID is not allowed")
	}
	return nil
}

func (sco *StudentCheckout) isStudentIDValid() *ValidationError {
	if sco.studentID == uuid.Nil {
		return NewValidationError("empty UUID in student ID is not allowed")
	}
	return nil
}

func (sco *StudentCheckout) isManavisCodeValid() *ValidationError {
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

func (sco *StudentCheckout) isCheckedOutAtValid() *ValidationError {
	if sco.checkedOutAt.IsZero() {
		return NewValidationError("empty checked out at is not allowed")
	}
	return nil
}

func (s *StudentCheckout) validate() *ValidationErrors {
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
	if err := s.isCheckedOutAtValid(); err != nil {
		errors = append(errors, err)
	}

	return validationErrorSliceToValidationErrors(errors)
}
