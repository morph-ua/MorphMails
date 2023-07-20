package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Connector holds the schema definition for the Connector entity.
type Connector struct {
	ent.Schema
}

// Annotations of the Connector.
func (Connector) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}

// Fields of the Connector.
func (Connector) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			StructTag(`json:"id" query:"id" form:"id"`).
			Unique(),
		field.
			String("name").
			StructTag(`json:"name" query:"name" form:"name"`),
		field.
			String("url").
			StructTag(`json:"url,omitempty" query:"url" form:"url"`).
			Unique(),
		field.
			String("secret").
			StructTag(`json:"secret,omitempty" query:"secret" form:"secret"`),
	}
}

// Edges of the Connector.
func (Connector) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("receivers", Receiver.Type),
	}
}
