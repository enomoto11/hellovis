package model

import (
	"errors"
	util "hellovis/api/utils"
	"strings"
)

type ValidationError struct {
	msg string
}

func (ve *ValidationError) Error() string {
	return ve.msg
}
func NewValidationError(msg string) *ValidationError {
	return &ValidationError{msg}
}

func IsValidationError(err error) bool {
	var validationErr *ValidationError
	return errors.As(err, &validationErr)
}

type ValidationErrors []*ValidationError

func (ves *ValidationErrors) Error() string {
	errMsgs := util.MapSlice(*ves, func(v *ValidationError) string {
		if v == nil {
			return ""
		}
		return v.Error()
	})
	return strings.Join(errMsgs, "")
}

func validationErrorSliceToValidationErrors(errors []*ValidationError) *ValidationErrors {
	if len(errors) != 0 {
		err := ValidationErrors(errors)
		return &err
	}
	return nil
}
