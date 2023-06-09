// Code generated by ent, DO NOT EDIT.

package ent

import (
	"hellovis/ent/schema"
	"hellovis/ent/student"
	"hellovis/ent/studentcheckin"
	"hellovis/ent/studentcheckout"
	"hellovis/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	studentMixin := schema.Student{}.Mixin()
	studentMixinFields0 := studentMixin[0].Fields()
	_ = studentMixinFields0
	studentMixinFields1 := studentMixin[1].Fields()
	_ = studentMixinFields1
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescCreatedAt is the schema descriptor for created_at field.
	studentDescCreatedAt := studentMixinFields0[0].Descriptor()
	// student.DefaultCreatedAt holds the default value on creation for the created_at field.
	student.DefaultCreatedAt = studentDescCreatedAt.Default.(func() time.Time)
	// studentDescUpdatedAt is the schema descriptor for updated_at field.
	studentDescUpdatedAt := studentMixinFields0[1].Descriptor()
	// student.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	student.DefaultUpdatedAt = studentDescUpdatedAt.Default.(func() time.Time)
	// student.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	student.UpdateDefaultUpdatedAt = studentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// studentDescLastName is the schema descriptor for last_name field.
	studentDescLastName := studentFields[0].Descriptor()
	// student.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	student.LastNameValidator = studentDescLastName.Validators[0].(func(string) error)
	// studentDescFirstName is the schema descriptor for first_name field.
	studentDescFirstName := studentFields[1].Descriptor()
	// student.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	student.FirstNameValidator = studentDescFirstName.Validators[0].(func(string) error)
	// studentDescGrade is the schema descriptor for grade field.
	studentDescGrade := studentFields[2].Descriptor()
	// student.GradeValidator is a validator for the "grade" field. It is called by the builders before save.
	student.GradeValidator = func() func(int) error {
		validators := studentDescGrade.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(grade int) error {
			for _, fn := range fns {
				if err := fn(grade); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// studentDescIsHighSchool is the schema descriptor for is_high_school field.
	studentDescIsHighSchool := studentFields[3].Descriptor()
	// student.DefaultIsHighSchool holds the default value on creation for the is_high_school field.
	student.DefaultIsHighSchool = studentDescIsHighSchool.Default.(bool)
	// studentDescManavisCode is the schema descriptor for manavis_code field.
	studentDescManavisCode := studentFields[4].Descriptor()
	// student.ManavisCodeValidator is a validator for the "manavis_code" field. It is called by the builders before save.
	student.ManavisCodeValidator = studentDescManavisCode.Validators[0].(func(string) error)
	// studentDescID is the schema descriptor for id field.
	studentDescID := studentMixinFields1[0].Descriptor()
	// student.DefaultID holds the default value on creation for the id field.
	student.DefaultID = studentDescID.Default.(func() uuid.UUID)
	studentcheckinMixin := schema.StudentCheckin{}.Mixin()
	studentcheckinMixinFields0 := studentcheckinMixin[0].Fields()
	_ = studentcheckinMixinFields0
	studentcheckinMixinFields1 := studentcheckinMixin[1].Fields()
	_ = studentcheckinMixinFields1
	studentcheckinFields := schema.StudentCheckin{}.Fields()
	_ = studentcheckinFields
	// studentcheckinDescCreatedAt is the schema descriptor for created_at field.
	studentcheckinDescCreatedAt := studentcheckinMixinFields0[0].Descriptor()
	// studentcheckin.DefaultCreatedAt holds the default value on creation for the created_at field.
	studentcheckin.DefaultCreatedAt = studentcheckinDescCreatedAt.Default.(func() time.Time)
	// studentcheckinDescUpdatedAt is the schema descriptor for updated_at field.
	studentcheckinDescUpdatedAt := studentcheckinMixinFields0[1].Descriptor()
	// studentcheckin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	studentcheckin.DefaultUpdatedAt = studentcheckinDescUpdatedAt.Default.(func() time.Time)
	// studentcheckin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	studentcheckin.UpdateDefaultUpdatedAt = studentcheckinDescUpdatedAt.UpdateDefault.(func() time.Time)
	// studentcheckinDescCheckinAt is the schema descriptor for checkin_at field.
	studentcheckinDescCheckinAt := studentcheckinFields[1].Descriptor()
	// studentcheckin.DefaultCheckinAt holds the default value on creation for the checkin_at field.
	studentcheckin.DefaultCheckinAt = studentcheckinDescCheckinAt.Default.(func() time.Time)
	// studentcheckinDescID is the schema descriptor for id field.
	studentcheckinDescID := studentcheckinMixinFields1[0].Descriptor()
	// studentcheckin.DefaultID holds the default value on creation for the id field.
	studentcheckin.DefaultID = studentcheckinDescID.Default.(func() uuid.UUID)
	studentcheckoutMixin := schema.StudentCheckout{}.Mixin()
	studentcheckoutMixinFields0 := studentcheckoutMixin[0].Fields()
	_ = studentcheckoutMixinFields0
	studentcheckoutMixinFields1 := studentcheckoutMixin[1].Fields()
	_ = studentcheckoutMixinFields1
	studentcheckoutFields := schema.StudentCheckout{}.Fields()
	_ = studentcheckoutFields
	// studentcheckoutDescCreatedAt is the schema descriptor for created_at field.
	studentcheckoutDescCreatedAt := studentcheckoutMixinFields0[0].Descriptor()
	// studentcheckout.DefaultCreatedAt holds the default value on creation for the created_at field.
	studentcheckout.DefaultCreatedAt = studentcheckoutDescCreatedAt.Default.(func() time.Time)
	// studentcheckoutDescUpdatedAt is the schema descriptor for updated_at field.
	studentcheckoutDescUpdatedAt := studentcheckoutMixinFields0[1].Descriptor()
	// studentcheckout.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	studentcheckout.DefaultUpdatedAt = studentcheckoutDescUpdatedAt.Default.(func() time.Time)
	// studentcheckout.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	studentcheckout.UpdateDefaultUpdatedAt = studentcheckoutDescUpdatedAt.UpdateDefault.(func() time.Time)
	// studentcheckoutDescCheckoutAt is the schema descriptor for checkout_at field.
	studentcheckoutDescCheckoutAt := studentcheckoutFields[1].Descriptor()
	// studentcheckout.DefaultCheckoutAt holds the default value on creation for the checkout_at field.
	studentcheckout.DefaultCheckoutAt = studentcheckoutDescCheckoutAt.Default.(func() time.Time)
	// studentcheckoutDescID is the schema descriptor for id field.
	studentcheckoutDescID := studentcheckoutMixinFields1[0].Descriptor()
	// studentcheckout.DefaultID holds the default value on creation for the id field.
	studentcheckout.DefaultID = studentcheckoutDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userMixinFields2 := userMixin[2].Fields()
	_ = userMixinFields2
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescSignInFailedCount is the schema descriptor for sign_in_failed_count field.
	userDescSignInFailedCount := userMixinFields2[0].Descriptor()
	// user.DefaultSignInFailedCount holds the default value on creation for the sign_in_failed_count field.
	user.DefaultSignInFailedCount = userDescSignInFailedCount.Default.(int8)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[0].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[1].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[2].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = func() func(string) error {
		validators := userDescPasswordHash.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password_hash string) error {
			for _, fn := range fns {
				if err := fn(password_hash); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields1[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
