package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Connector holds the schema definition for the Connector entity.
type Connector struct {
	ent.Schema
}

// Fields of the Connector.
func (Connector) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			StructTag(`json:"id" query:"id" form:"id"`),
		field.
			String("name").
			StructTag(`json:"name" query:"name" form:"name"`),
		field.
			String("url").
			StructTag(`json:"url" query:"url" form:"url"`),
		field.
			String("secret").
			StructTag(`json:"secret" query:"secret" form:"secret"`),
	}
}

// Edges of the Connector.
func (Connector) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("receivers", Receiver.Type),
	}
}
