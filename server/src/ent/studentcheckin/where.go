// Code generated by ent, DO NOT EDIT.

package studentcheckin

import (
	"hellovis/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldUpdatedAt, v))
}

// StudentID applies equality check predicate on the "student_id" field. It's identical to StudentIDEQ.
func StudentID(v uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldStudentID, v))
}

// CheckinAt applies equality check predicate on the "checkin_at" field. It's identical to CheckinAtEQ.
func CheckinAt(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldCheckinAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLTE(FieldUpdatedAt, v))
}

// StudentIDEQ applies the EQ predicate on the "student_id" field.
func StudentIDEQ(v uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldStudentID, v))
}

// StudentIDNEQ applies the NEQ predicate on the "student_id" field.
func StudentIDNEQ(v uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNEQ(FieldStudentID, v))
}

// StudentIDIn applies the In predicate on the "student_id" field.
func StudentIDIn(vs ...uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldIn(FieldStudentID, vs...))
}

// StudentIDNotIn applies the NotIn predicate on the "student_id" field.
func StudentIDNotIn(vs ...uuid.UUID) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNotIn(FieldStudentID, vs...))
}

// CheckinAtEQ applies the EQ predicate on the "checkin_at" field.
func CheckinAtEQ(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldEQ(FieldCheckinAt, v))
}

// CheckinAtNEQ applies the NEQ predicate on the "checkin_at" field.
func CheckinAtNEQ(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNEQ(FieldCheckinAt, v))
}

// CheckinAtIn applies the In predicate on the "checkin_at" field.
func CheckinAtIn(vs ...time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldIn(FieldCheckinAt, vs...))
}

// CheckinAtNotIn applies the NotIn predicate on the "checkin_at" field.
func CheckinAtNotIn(vs ...time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldNotIn(FieldCheckinAt, vs...))
}

// CheckinAtGT applies the GT predicate on the "checkin_at" field.
func CheckinAtGT(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGT(FieldCheckinAt, v))
}

// CheckinAtGTE applies the GTE predicate on the "checkin_at" field.
func CheckinAtGTE(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldGTE(FieldCheckinAt, v))
}

// CheckinAtLT applies the LT predicate on the "checkin_at" field.
func CheckinAtLT(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLT(FieldCheckinAt, v))
}

// CheckinAtLTE applies the LTE predicate on the "checkin_at" field.
func CheckinAtLTE(v time.Time) predicate.StudentCheckin {
	return predicate.StudentCheckin(sql.FieldLTE(FieldCheckinAt, v))
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.StudentCheckin {
	return predicate.StudentCheckin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.StudentCheckin {
	return predicate.StudentCheckin(func(s *sql.Selector) {
		step := newStudentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StudentCheckin) predicate.StudentCheckin {
	return predicate.StudentCheckin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StudentCheckin) predicate.StudentCheckin {
	return predicate.StudentCheckin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StudentCheckin) predicate.StudentCheckin {
	return predicate.StudentCheckin(func(s *sql.Selector) {
		p(s.Not())
	})
}
