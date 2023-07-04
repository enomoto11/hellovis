package schema

import (
	"hellovis/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("first_name").NotEmpty(),
		field.String("password_hash").MaxLen(191).NotEmpty(),
		field.String("email").MaxLen(191).NotEmpty().Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.UUIDMixin{},
		mixin.AccountLockMixin{},
	}
}
