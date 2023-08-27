package schema

import (
	"hellovis/ent/schema/mixin"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentCheckin holds the schema definition for the StudentCheckin entity.
type StudentCheckin struct {
	ent.Schema
}

// Fields of the StudentCheckin.
func (StudentCheckin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("student_id", uuid.UUID{}),
		field.Time("checkin_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime(6)",
			}),
	}
}

// Edges of the StudentCheckin.
func (StudentCheckin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).
			Ref("checkins").
			Unique().
			Field("student_id").
			Required(),
	}
}

// Mixin of the StudentCheckin.
func (StudentCheckin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
	}
}
