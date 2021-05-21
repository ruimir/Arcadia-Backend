package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Datafile holds the schema definition for the Datafile entity.
type Datafile struct {
	ent.Schema
}

// Fields of the Datafile.
func (Datafile) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Datafile.
func (Datafile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("header", Header.Type).Unique().StructTag(`xml:"header"`),
		edge.To("games", Game.Type).StructTag(`xml:"game"`),
	}
}
