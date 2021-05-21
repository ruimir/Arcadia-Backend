// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Backend/ent/datafile"
	"Backend/ent/game"
	"Backend/ent/release"
	"Backend/ent/rom"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GameCreate is the builder for creating a Game entity.
type GameCreate struct {
	config
	mutation *GameMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GameCreate) SetName(s string) *GameCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetCloneof sets the "cloneof" field.
func (gc *GameCreate) SetCloneof(s string) *GameCreate {
	gc.mutation.SetCloneof(s)
	return gc
}

// SetDescription sets the "description" field.
func (gc *GameCreate) SetDescription(s string) *GameCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetDatafileID sets the "datafile" edge to the Datafile entity by ID.
func (gc *GameCreate) SetDatafileID(id int) *GameCreate {
	gc.mutation.SetDatafileID(id)
	return gc
}

// SetNillableDatafileID sets the "datafile" edge to the Datafile entity by ID if the given value is not nil.
func (gc *GameCreate) SetNillableDatafileID(id *int) *GameCreate {
	if id != nil {
		gc = gc.SetDatafileID(*id)
	}
	return gc
}

// SetDatafile sets the "datafile" edge to the Datafile entity.
func (gc *GameCreate) SetDatafile(d *Datafile) *GameCreate {
	return gc.SetDatafileID(d.ID)
}

// AddReleaseIDs adds the "releases" edge to the Release entity by IDs.
func (gc *GameCreate) AddReleaseIDs(ids ...int) *GameCreate {
	gc.mutation.AddReleaseIDs(ids...)
	return gc
}

// AddReleases adds the "releases" edges to the Release entity.
func (gc *GameCreate) AddReleases(r ...*Release) *GameCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return gc.AddReleaseIDs(ids...)
}

// SetRomID sets the "rom" edge to the Rom entity by ID.
func (gc *GameCreate) SetRomID(id int) *GameCreate {
	gc.mutation.SetRomID(id)
	return gc
}

// SetNillableRomID sets the "rom" edge to the Rom entity by ID if the given value is not nil.
func (gc *GameCreate) SetNillableRomID(id *int) *GameCreate {
	if id != nil {
		gc = gc.SetRomID(*id)
	}
	return gc
}

// SetRom sets the "rom" edge to the Rom entity.
func (gc *GameCreate) SetRom(r *Rom) *GameCreate {
	return gc.SetRomID(r.ID)
}

// Mutation returns the GameMutation object of the builder.
func (gc *GameCreate) Mutation() *GameMutation {
	return gc.mutation
}

// Save creates the Game in the database.
func (gc *GameCreate) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GameCreate) SaveX(ctx context.Context) *Game {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (gc *GameCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := gc.mutation.Cloneof(); !ok {
		return &ValidationError{Name: "cloneof", err: errors.New("ent: missing required field \"cloneof\"")}
	}
	if _, ok := gc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	return nil
}

func (gc *GameCreate) sqlSave(ctx context.Context) (*Game, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GameCreate) createSpec() (*Game, *sqlgraph.CreateSpec) {
	var (
		_node = &Game{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: game.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.Cloneof(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldCloneof,
		})
		_node.Cloneof = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := gc.mutation.DatafileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.DatafileTable,
			Columns: []string{game.DatafileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: datafile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.datafile_games = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.ReleasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   game.ReleasesTable,
			Columns: []string{game.ReleasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: release.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.RomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   game.RomTable,
			Columns: []string{game.RomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GameCreateBulk is the builder for creating many Game entities in bulk.
type GameCreateBulk struct {
	config
	builders []*GameCreate
}

// Save creates the Game entities in the database.
func (gcb *GameCreateBulk) Save(ctx context.Context) ([]*Game, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Game, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GameMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GameCreateBulk) SaveX(ctx context.Context) []*Game {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
