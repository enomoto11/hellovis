package schema

import (
	"hellovis/ent/schema/mixin"

	"entgo.io/ent"
)

// StudentCheckout holds the schema definition for the StudentCheckout entity.
type StudentCheckout struct {
	ent.Schema
}

// Fields of the StudentCheckout.
func (StudentCheckout) Fields() []ent.Field {
	return nil
}

// Edges of the StudentCheckout.
func (StudentCheckout) Edges() []ent.Edge {
	return nil
}

// Mixin of the Student.
func (StudentCheckout) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
	}
}
