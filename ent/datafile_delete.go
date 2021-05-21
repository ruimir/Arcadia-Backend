// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Backend/ent/datafile"
	"Backend/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DatafileDelete is the builder for deleting a Datafile entity.
type DatafileDelete struct {
	config
	hooks    []Hook
	mutation *DatafileMutation
}

// Where adds a new predicate to the DatafileDelete builder.
func (dd *DatafileDelete) Where(ps ...predicate.Datafile) *DatafileDelete {
	dd.mutation.predicates = append(dd.mutation.predicates, ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DatafileDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dd.hooks) == 0 {
		affected, err = dd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatafileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dd.mutation = mutation
			affected, err = dd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dd.hooks) - 1; i >= 0; i-- {
			mut = dd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DatafileDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DatafileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: datafile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: datafile.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
}

// DatafileDeleteOne is the builder for deleting a single Datafile entity.
type DatafileDeleteOne struct {
	dd *DatafileDelete
}

// Exec executes the deletion query.
func (ddo *DatafileDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{datafile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DatafileDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
