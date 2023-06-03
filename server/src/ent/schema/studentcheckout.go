package schema

import (
	"hellovis/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentCheckout holds the schema definition for the StudentCheckout entity.
type StudentCheckout struct {
	ent.Schema
}

// Fields of the StudentCheckout.
func (StudentCheckout) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("student_id", uuid.UUID{}),
	}
}

// Edges of the StudentCheckout.
func (StudentCheckout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).
			Ref("checkouts").
			Unique().
			Field("student_id").
			Required(),
	}
}

// Mixin of the StudentCheckout.
func (StudentCheckout) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
	}
}
