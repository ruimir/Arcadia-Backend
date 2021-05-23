package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rom holds the schema definition for the Rom entity.
type Rom struct {
	ent.Schema
}

// Fields of the Rom.
func (Rom) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("size"),
		field.String("crc"),
		field.String("md5"),
		field.String("sha1"),
		field.String("status"),
	}
}

// Edges of the Rom.
func (Rom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).Ref("rom").Unique().Required(),
		edge.From("file", File.Type).Ref("rom"),
	}
}
