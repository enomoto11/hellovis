package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type AccountLockMixin struct {
	mixin.Schema
}

func (AccountLockMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("sign_in_failed_count").Default(0),
		field.Time("account_locked_until").Optional().Nillable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}
