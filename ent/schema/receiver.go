package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Receiver holds the schema definition for the Receiver entity.
type Receiver struct {
	ent.Schema
}

// Fields of the Receiver.
func (Receiver) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
	}
}

// Edges of the Receiver.
func (Receiver) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("receivers").
			Unique().
			Required(),
		edge.From("connector", Connector.Type).
			Ref("receivers").
			Unique().
			Required(),
	}
}
