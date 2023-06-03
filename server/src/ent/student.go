// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hellovis/ent/student"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// Grade holds the value of the "grade" field.
	Grade int16 `json:"grade,omitempty"`
	// IsHighSchool holds the value of the "is_high_school" field.
	IsHighSchool bool `json:"is_high_school,omitempty"`
	// ManavisCode holds the value of the "manavis_code" field.
	ManavisCode string `json:"manavis_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges        StudentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// Checkins holds the value of the checkins edge.
	Checkins []*StudentCheckin `json:"checkins,omitempty"`
	// Checkouts holds the value of the checkouts edge.
	Checkouts []*StudentCheckout `json:"checkouts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CheckinsOrErr returns the Checkins value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) CheckinsOrErr() ([]*StudentCheckin, error) {
	if e.loadedTypes[0] {
		return e.Checkins, nil
	}
	return nil, &NotLoadedError{edge: "checkins"}
}

// CheckoutsOrErr returns the Checkouts value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) CheckoutsOrErr() ([]*StudentCheckout, error) {
	if e.loadedTypes[1] {
		return e.Checkouts, nil
	}
	return nil, &NotLoadedError{edge: "checkouts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldIsHighSchool:
			values[i] = new(sql.NullBool)
		case student.FieldGrade:
			values[i] = new(sql.NullInt64)
		case student.FieldLastName, student.FieldFirstName, student.FieldManavisCode:
			values[i] = new(sql.NullString)
		case student.FieldCreatedAt, student.FieldUpdatedAt, student.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case student.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case student.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case student.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case student.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case student.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				s.LastName = value.String
			}
		case student.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				s.FirstName = value.String
			}
		case student.FieldGrade:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field grade", values[i])
			} else if value.Valid {
				s.Grade = int16(value.Int64)
			}
		case student.FieldIsHighSchool:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_high_school", values[i])
			} else if value.Valid {
				s.IsHighSchool = value.Bool
			}
		case student.FieldManavisCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manavis_code", values[i])
			} else if value.Valid {
				s.ManavisCode = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Student.
// This includes values selected through modifiers, order, etc.
func (s *Student) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryCheckins queries the "checkins" edge of the Student entity.
func (s *Student) QueryCheckins() *StudentCheckinQuery {
	return NewStudentClient(s.config).QueryCheckins(s)
}

// QueryCheckouts queries the "checkouts" edge of the Student entity.
func (s *Student) QueryCheckouts() *StudentCheckoutQuery {
	return NewStudentClient(s.config).QueryCheckouts(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return NewStudentClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(s.LastName)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(s.FirstName)
	builder.WriteString(", ")
	builder.WriteString("grade=")
	builder.WriteString(fmt.Sprintf("%v", s.Grade))
	builder.WriteString(", ")
	builder.WriteString("is_high_school=")
	builder.WriteString(fmt.Sprintf("%v", s.IsHighSchool))
	builder.WriteString(", ")
	builder.WriteString("manavis_code=")
	builder.WriteString(s.ManavisCode)
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student
