package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/bodokaiser/entgo-time-mixin/ent/schema/hook"
	"github.com/bodokaiser/entgo-time-mixin/ent/schema/mixin"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("slug"),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}

// Hooks of the Pet.
func (Pet) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.GenerateSlug,
	}
}

// Mixins of the Pet.
func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
