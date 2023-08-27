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

// StudentCheckout holds the schema definition for the StudentCheckout entity.
type StudentCheckout struct {
	ent.Schema
}

// Fields of the StudentCheckout.
func (StudentCheckout) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("student_id", uuid.UUID{}),
		field.Time("checkout_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime(6)",
			}),
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
