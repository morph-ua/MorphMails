// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"helium/ent/receiver"
	"helium/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetEmails sets the "emails" field.
func (uc *UserCreate) SetEmails(s []string) *UserCreate {
	uc.mutation.SetEmails(s)
	return uc
}

// SetForward sets the "forward" field.
func (uc *UserCreate) SetForward(b bool) *UserCreate {
	uc.mutation.SetForward(b)
	return uc
}

// SetNillableForward sets the "forward" field if the given value is not nil.
func (uc *UserCreate) SetNillableForward(b *bool) *UserCreate {
	if b != nil {
		uc.SetForward(*b)
	}
	return uc
}

// SetPaid sets the "paid" field.
func (uc *UserCreate) SetPaid(b bool) *UserCreate {
	uc.mutation.SetPaid(b)
	return uc
}

// SetNillablePaid sets the "paid" field if the given value is not nil.
func (uc *UserCreate) SetNillablePaid(b *bool) *UserCreate {
	if b != nil {
		uc.SetPaid(*b)
	}
	return uc
}

// SetCounter sets the "counter" field.
func (uc *UserCreate) SetCounter(i int8) *UserCreate {
	uc.mutation.SetCounter(i)
	return uc
}

// SetNillableCounter sets the "counter" field if the given value is not nil.
func (uc *UserCreate) SetNillableCounter(i *int8) *UserCreate {
	if i != nil {
		uc.SetCounter(*i)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// AddReceiverIDs adds the "receivers" edge to the Receiver entity by IDs.
func (uc *UserCreate) AddReceiverIDs(ids ...string) *UserCreate {
	uc.mutation.AddReceiverIDs(ids...)
	return uc
}

// AddReceivers adds the "receivers" edges to the Receiver entity.
func (uc *UserCreate) AddReceivers(r ...*Receiver) *UserCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddReceiverIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Forward(); !ok {
		v := user.DefaultForward
		uc.mutation.SetForward(v)
	}
	if _, ok := uc.mutation.Paid(); !ok {
		v := user.DefaultPaid
		uc.mutation.SetPaid(v)
	}
	if _, ok := uc.mutation.Counter(); !ok {
		v := user.DefaultCounter
		uc.mutation.SetCounter(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Forward(); !ok {
		return &ValidationError{Name: "forward", err: errors.New(`ent: missing required field "User.forward"`)}
	}
	if _, ok := uc.mutation.Paid(); !ok {
		return &ValidationError{Name: "paid", err: errors.New(`ent: missing required field "User.paid"`)}
	}
	if _, ok := uc.mutation.Counter(); !ok {
		return &ValidationError{Name: "counter", err: errors.New(`ent: missing required field "User.counter"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.Emails(); ok {
		_spec.SetField(user.FieldEmails, field.TypeJSON, value)
		_node.Emails = value
	}
	if value, ok := uc.mutation.Forward(); ok {
		_spec.SetField(user.FieldForward, field.TypeBool, value)
		_node.Forward = value
	}
	if value, ok := uc.mutation.Paid(); ok {
		_spec.SetField(user.FieldPaid, field.TypeBool, value)
		_node.Paid = value
	}
	if value, ok := uc.mutation.Counter(); ok {
		_spec.SetField(user.FieldCounter, field.TypeInt8, value)
		_node.Counter = value
	}
	if nodes := uc.mutation.ReceiversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReceiversTable,
			Columns: []string{user.ReceiversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receiver.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}