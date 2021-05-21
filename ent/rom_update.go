// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Backend/ent/game"
	"Backend/ent/predicate"
	"Backend/ent/rom"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RomUpdate is the builder for updating Rom entities.
type RomUpdate struct {
	config
	hooks    []Hook
	mutation *RomMutation
}

// Where adds a new predicate for the RomUpdate builder.
func (ru *RomUpdate) Where(ps ...predicate.Rom) *RomUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RomUpdate) SetName(s string) *RomUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetSize sets the "size" field.
func (ru *RomUpdate) SetSize(s string) *RomUpdate {
	ru.mutation.SetSize(s)
	return ru
}

// SetCrc sets the "crc" field.
func (ru *RomUpdate) SetCrc(s string) *RomUpdate {
	ru.mutation.SetCrc(s)
	return ru
}

// SetMd5 sets the "md5" field.
func (ru *RomUpdate) SetMd5(s string) *RomUpdate {
	ru.mutation.SetMd5(s)
	return ru
}

// SetSha1 sets the "sha1" field.
func (ru *RomUpdate) SetSha1(s string) *RomUpdate {
	ru.mutation.SetSha1(s)
	return ru
}

// SetStatus sets the "status" field.
func (ru *RomUpdate) SetStatus(s string) *RomUpdate {
	ru.mutation.SetStatus(s)
	return ru
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (ru *RomUpdate) SetGameID(id int) *RomUpdate {
	ru.mutation.SetGameID(id)
	return ru
}

// SetGame sets the "game" edge to the Game entity.
func (ru *RomUpdate) SetGame(g *Game) *RomUpdate {
	return ru.SetGameID(g.ID)
}

// Mutation returns the RomMutation object of the builder.
func (ru *RomUpdate) Mutation() *RomMutation {
	return ru.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (ru *RomUpdate) ClearGame() *RomUpdate {
	ru.mutation.ClearGame()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RomUpdate) check() error {
	if _, ok := ru.mutation.GameID(); ru.mutation.GameCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"game\"")
	}
	return nil
}

func (ru *RomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rom.Table,
			Columns: rom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rom.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldName,
		})
	}
	if value, ok := ru.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldSize,
		})
	}
	if value, ok := ru.mutation.Crc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldCrc,
		})
	}
	if value, ok := ru.mutation.Md5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldMd5,
		})
	}
	if value, ok := ru.mutation.Sha1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldSha1,
		})
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldStatus,
		})
	}
	if ru.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rom.GameTable,
			Columns: []string{rom.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rom.GameTable,
			Columns: []string{rom.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RomUpdateOne is the builder for updating a single Rom entity.
type RomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RomMutation
}

// SetName sets the "name" field.
func (ruo *RomUpdateOne) SetName(s string) *RomUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetSize sets the "size" field.
func (ruo *RomUpdateOne) SetSize(s string) *RomUpdateOne {
	ruo.mutation.SetSize(s)
	return ruo
}

// SetCrc sets the "crc" field.
func (ruo *RomUpdateOne) SetCrc(s string) *RomUpdateOne {
	ruo.mutation.SetCrc(s)
	return ruo
}

// SetMd5 sets the "md5" field.
func (ruo *RomUpdateOne) SetMd5(s string) *RomUpdateOne {
	ruo.mutation.SetMd5(s)
	return ruo
}

// SetSha1 sets the "sha1" field.
func (ruo *RomUpdateOne) SetSha1(s string) *RomUpdateOne {
	ruo.mutation.SetSha1(s)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RomUpdateOne) SetStatus(s string) *RomUpdateOne {
	ruo.mutation.SetStatus(s)
	return ruo
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (ruo *RomUpdateOne) SetGameID(id int) *RomUpdateOne {
	ruo.mutation.SetGameID(id)
	return ruo
}

// SetGame sets the "game" edge to the Game entity.
func (ruo *RomUpdateOne) SetGame(g *Game) *RomUpdateOne {
	return ruo.SetGameID(g.ID)
}

// Mutation returns the RomMutation object of the builder.
func (ruo *RomUpdateOne) Mutation() *RomMutation {
	return ruo.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (ruo *RomUpdateOne) ClearGame() *RomUpdateOne {
	ruo.mutation.ClearGame()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RomUpdateOne) Select(field string, fields ...string) *RomUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Rom entity.
func (ruo *RomUpdateOne) Save(ctx context.Context) (*Rom, error) {
	var (
		err  error
		node *Rom
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RomUpdateOne) SaveX(ctx context.Context) *Rom {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RomUpdateOne) check() error {
	if _, ok := ruo.mutation.GameID(); ruo.mutation.GameCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"game\"")
	}
	return nil
}

func (ruo *RomUpdateOne) sqlSave(ctx context.Context) (_node *Rom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rom.Table,
			Columns: rom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rom.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Rom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rom.FieldID)
		for _, f := range fields {
			if !rom.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldName,
		})
	}
	if value, ok := ruo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldSize,
		})
	}
	if value, ok := ruo.mutation.Crc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldCrc,
		})
	}
	if value, ok := ruo.mutation.Md5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldMd5,
		})
	}
	if value, ok := ruo.mutation.Sha1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldSha1,
		})
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rom.FieldStatus,
		})
	}
	if ruo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rom.GameTable,
			Columns: []string{rom.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rom.GameTable,
			Columns: []string{rom.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rom{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}