package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Header holds the schema definition for the Header entity.
type Header struct {
	ent.Schema
}

// Fields of the Header.
func (Header) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`xml:"name"`),
		field.String("description"),
		field.String("version"),
		field.String("date"),
		field.String("author"),
		field.String("url"),
	}
}

// Edges of the Header.
func (Header) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("datafile", Datafile.Type).Ref("header").Unique().Required(),
	}
}
