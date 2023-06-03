package schema

import (
	"hellovis/ent/schema/mixin"

	"entgo.io/ent"
)

// StudentCheckin holds the schema definition for the StudentCheckin entity.
type StudentCheckin struct {
	ent.Schema
}

// Fields of the StudentCheckin.
func (StudentCheckin) Fields() []ent.Field {
	return nil
}

// Edges of the StudentCheckin.
func (StudentCheckin) Edges() []ent.Edge {
	return nil
}

// Mixin of the Student.
func (StudentCheckin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
	}
}
