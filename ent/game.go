// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Backend/ent/datafile"
	"Backend/ent/game"
	"Backend/ent/rom"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Game is the model entity for the Game schema.
type Game struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Cloneof holds the value of the "cloneof" field.
	Cloneof string `json:"cloneof,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameQuery when eager-loading is set.
	Edges          GameEdges `json:"edges"`
	datafile_games *int
}

// GameEdges holds the relations/edges for other nodes in the graph.
type GameEdges struct {
	// Datafile holds the value of the datafile edge.
	Datafile *Datafile `json:"datafile,omitempty"`
	// Releases holds the value of the releases edge.
	Releases []*Release `json:"releases,omitempty" xml:"release"`
	// Rom holds the value of the rom edge.
	Rom *Rom `json:"rom,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DatafileOrErr returns the Datafile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) DatafileOrErr() (*Datafile, error) {
	if e.loadedTypes[0] {
		if e.Datafile == nil {
			// The edge datafile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: datafile.Label}
		}
		return e.Datafile, nil
	}
	return nil, &NotLoadedError{edge: "datafile"}
}

// ReleasesOrErr returns the Releases value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) ReleasesOrErr() ([]*Release, error) {
	if e.loadedTypes[1] {
		return e.Releases, nil
	}
	return nil, &NotLoadedError{edge: "releases"}
}

// RomOrErr returns the Rom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) RomOrErr() (*Rom, error) {
	if e.loadedTypes[2] {
		if e.Rom == nil {
			// The edge rom was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: rom.Label}
		}
		return e.Rom, nil
	}
	return nil, &NotLoadedError{edge: "rom"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Game) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			values[i] = new(sql.NullInt64)
		case game.FieldName, game.FieldCloneof, game.FieldDescription:
			values[i] = new(sql.NullString)
		case game.ForeignKeys[0]: // datafile_games
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Game", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Game fields.
func (ga *Game) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case game.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ga.Name = value.String
			}
		case game.FieldCloneof:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cloneof", values[i])
			} else if value.Valid {
				ga.Cloneof = value.String
			}
		case game.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ga.Description = value.String
			}
		case game.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field datafile_games", value)
			} else if value.Valid {
				ga.datafile_games = new(int)
				*ga.datafile_games = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDatafile queries the "datafile" edge of the Game entity.
func (ga *Game) QueryDatafile() *DatafileQuery {
	return (&GameClient{config: ga.config}).QueryDatafile(ga)
}

// QueryReleases queries the "releases" edge of the Game entity.
func (ga *Game) QueryReleases() *ReleaseQuery {
	return (&GameClient{config: ga.config}).QueryReleases(ga)
}

// QueryRom queries the "rom" edge of the Game entity.
func (ga *Game) QueryRom() *RomQuery {
	return (&GameClient{config: ga.config}).QueryRom(ga)
}

// Update returns a builder for updating this Game.
// Note that you need to call Game.Unwrap() before calling this method if this Game
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Game) Update() *GameUpdateOne {
	return (&GameClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Game entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Game) Unwrap() *Game {
	tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Game is not a transactional entity")
	}
	ga.config.driver = tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Game) String() string {
	var builder strings.Builder
	builder.WriteString("Game(")
	builder.WriteString(fmt.Sprintf("id=%v", ga.ID))
	builder.WriteString(", name=")
	builder.WriteString(ga.Name)
	builder.WriteString(", cloneof=")
	builder.WriteString(ga.Cloneof)
	builder.WriteString(", description=")
	builder.WriteString(ga.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Games is a parsable slice of Game.
type Games []*Game

func (ga Games) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
