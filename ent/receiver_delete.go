// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"helium/ent/predicate"
	"helium/ent/receiver"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReceiverDelete is the builder for deleting a Receiver entity.
type ReceiverDelete struct {
	config
	hooks    []Hook
	mutation *ReceiverMutation
}

// Where appends a list predicates to the ReceiverDelete builder.
func (rd *ReceiverDelete) Where(ps ...predicate.Receiver) *ReceiverDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ReceiverDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ReceiverDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ReceiverDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(receiver.Table, sqlgraph.NewFieldSpec(receiver.FieldID, field.TypeString))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// ReceiverDeleteOne is the builder for deleting a single Receiver entity.
type ReceiverDeleteOne struct {
	rd *ReceiverDelete
}

// Where appends a list predicates to the ReceiverDelete builder.
func (rdo *ReceiverDeleteOne) Where(ps ...predicate.Receiver) *ReceiverDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *ReceiverDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{receiver.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ReceiverDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}