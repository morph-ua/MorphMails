package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"net/mail"
	"time"
)

// Letter holds the schema definition for the Letter entity.
type Letter struct {
	ent.Schema
}

// Fields of the Letter.
func (Letter) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").Unique(),
		field.
			String("html"),
		field.
			String("from").
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				return err
			}),
		field.
			String("to").
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				return err
			}),
		field.
			Time("created_at").
			Default(time.Now()).
			Immutable().
			Comment("Save time to delete after 3 days"),
	}
}

// Edges of the Letter.
func (Letter) Edges() []ent.Edge {
	return nil
}
