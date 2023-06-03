package schema

import (
	"hellovis/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("first_name").NotEmpty(),
		field.Int16("grade").Min(1).Max(3),
		field.String("manavis_code").NotEmpty().Unique(),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checkins", StudentCheckin.Type),
		edge.To("checkouts", StudentCheckout.Type),
	}
}

// Mixin of the Student.
func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
	}
}
