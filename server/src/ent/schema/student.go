package schema

import (
	"hellovis/ent/schema/mixin"

	"entgo.io/ent"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return nil
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}

// Mixin of the Student.
func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
	}
}
