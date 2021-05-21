package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("cloneof"),
		field.String("description"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("datafile", Datafile.Type).Ref("games").Unique(),
		edge.To("releases", Release.Type).StructTag(`xml:"release"`),
		edge.To("rom", Rom.Type).Unique(),
	}
}
