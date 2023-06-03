package model

import (
	"strconv"

	"github.com/google/uuid"
)

type Student struct {
	id             uuid.UUID `json:"id"`
	firstName      string    `json:"first_name"`
	lastName       string    `json:"last_name"`
	grade          int       `json:"grade"`
	manavisCode    string    `json:"manavis_code"`
	isInHighSchool bool      `json:"is_high_school"`
}

type NewStudentOption func(*Student)

func NewStudent(opts ...NewStudentOption) (*Student, error) {
	s := &Student{}
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

func NewStudentID(id uuid.UUID) NewStudentOption {
	return func(s *Student) {
		s.id = id
	}
}

func NewStudentFirstName(firstName string) NewStudentOption {
	return func(s *Student) {
		s.firstName = firstName
	}
}

func NewStudentLastName(lastName string) NewStudentOption {
	return func(s *Student) {
		s.lastName = lastName
	}
}

func NewStudentGrade(grade int) NewStudentOption {
	return func(s *Student) {
		s.grade = grade
	}
}

func NewStudentManavisCode(manavisCode string) NewStudentOption {
	return func(s *Student) {
		s.manavisCode = manavisCode
	}
}

func NewStudentIsInHighSchool(isInHighSchool bool) NewStudentOption {
	return func(s *Student) {
		s.isInHighSchool = isInHighSchool
	}
}

func (s *Student) GetID() uuid.UUID {
	return s.id
}

func (s *Student) GetFirstName() string {
	return s.firstName
}

func (s *Student) GetLastName() string {
	return s.lastName
}

func (s *Student) GetGrade() int {
	return s.grade
}

func (s *Student) GetManavisCode() string {
	return s.manavisCode
}

func (s *Student) GetIsInHighSchool() bool {
	return s.isInHighSchool
}

func (s *Student) isIDValid() *ValidationError {
	if s.id == uuid.Nil {
		return NewValidationError("empty UUID in member ID is not allowed")
	}
	return nil
}

func (s *Student) isFirstNameValid() *ValidationError {
	if s.firstName == "" {
		return NewValidationError("empty first name is not allowed")
	}
	return nil
}

func (s *Student) isLastNameValid() *ValidationError {
	if s.lastName == "" {
		return NewValidationError("empty last name is not allowed")
	}
	return nil
}

func (s *Student) isGradeValid() *ValidationError {
	if s.grade != 1 && s.grade != 2 && s.grade != 3 {
		return NewValidationError("grade must be 1, 2 or 3")
	}
	return nil
}

func (s *Student) isManavisCodeValid() *ValidationError {
	if s.manavisCode == "" {
		return NewValidationError("empty manavis code is not allowed")
	}
	num, err := strconv.Atoi(s.manavisCode)
	if err != nil {
		return NewValidationError("manavis code must be a number")
	}
	if num <= 0 {
		return NewValidationError("manavis code must be positive")
	}
	return nil
}

func (s *Student) validate() *ValidationErrors {
	var errors []*ValidationError
	if err := s.isIDValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isFirstNameValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isLastNameValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isGradeValid(); err != nil {
		errors = append(errors, err)
	}
	if err := s.isManavisCodeValid(); err != nil {
		errors = append(errors, err)
	}
	return validationErrorSliceToValidationErrors(errors)
}
