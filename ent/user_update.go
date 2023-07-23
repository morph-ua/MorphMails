// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"helium/ent/predicate"
	"helium/ent/receiver"
	"helium/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmails sets the "emails" field.
func (uu *UserUpdate) SetEmails(s []string) *UserUpdate {
	uu.mutation.SetEmails(s)
	return uu
}

// AppendEmails appends s to the "emails" field.
func (uu *UserUpdate) AppendEmails(s []string) *UserUpdate {
	uu.mutation.AppendEmails(s)
	return uu
}

// ClearEmails clears the value of the "emails" field.
func (uu *UserUpdate) ClearEmails() *UserUpdate {
	uu.mutation.ClearEmails()
	return uu
}

// SetForward sets the "forward" field.
func (uu *UserUpdate) SetForward(b bool) *UserUpdate {
	uu.mutation.SetForward(b)
	return uu
}

// SetNillableForward sets the "forward" field if the given value is not nil.
func (uu *UserUpdate) SetNillableForward(b *bool) *UserUpdate {
	if b != nil {
		uu.SetForward(*b)
	}
	return uu
}

// SetPaid sets the "paid" field.
func (uu *UserUpdate) SetPaid(b bool) *UserUpdate {
	uu.mutation.SetPaid(b)
	return uu
}

// SetNillablePaid sets the "paid" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePaid(b *bool) *UserUpdate {
	if b != nil {
		uu.SetPaid(*b)
	}
	return uu
}

// SetCounter sets the "counter" field.
func (uu *UserUpdate) SetCounter(i int8) *UserUpdate {
	uu.mutation.ResetCounter()
	uu.mutation.SetCounter(i)
	return uu
}

// SetNillableCounter sets the "counter" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCounter(i *int8) *UserUpdate {
	if i != nil {
		uu.SetCounter(*i)
	}
	return uu
}

// AddCounter adds i to the "counter" field.
func (uu *UserUpdate) AddCounter(i int8) *UserUpdate {
	uu.mutation.AddCounter(i)
	return uu
}

// AddReceiverIDs adds the "receivers" edge to the Receiver entity by IDs.
func (uu *UserUpdate) AddReceiverIDs(ids ...string) *UserUpdate {
	uu.mutation.AddReceiverIDs(ids...)
	return uu
}

// AddReceivers adds the "receivers" edges to the Receiver entity.
func (uu *UserUpdate) AddReceivers(r ...*Receiver) *UserUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddReceiverIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearReceivers clears all "receivers" edges to the Receiver entity.
func (uu *UserUpdate) ClearReceivers() *UserUpdate {
	uu.mutation.ClearReceivers()
	return uu
}

// RemoveReceiverIDs removes the "receivers" edge to Receiver entities by IDs.
func (uu *UserUpdate) RemoveReceiverIDs(ids ...string) *UserUpdate {
	uu.mutation.RemoveReceiverIDs(ids...)
	return uu
}

// RemoveReceivers removes "receivers" edges to Receiver entities.
func (uu *UserUpdate) RemoveReceivers(r ...*Receiver) *UserUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveReceiverIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Emails(); ok {
		_spec.SetField(user.FieldEmails, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedEmails(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldEmails, value)
		})
	}
	if uu.mutation.EmailsCleared() {
		_spec.ClearField(user.FieldEmails, field.TypeJSON)
	}
	if value, ok := uu.mutation.Forward(); ok {
		_spec.SetField(user.FieldForward, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Paid(); ok {
		_spec.SetField(user.FieldPaid, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Counter(); ok {
		_spec.SetField(user.FieldCounter, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedCounter(); ok {
		_spec.AddField(user.FieldCounter, field.TypeInt8, value)
	}
	if uu.mutation.CreatedAtCleared() {
		_spec.ClearField(user.FieldCreatedAt, field.TypeTime)
	}
	if uu.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !uu.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ReceiversIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmails sets the "emails" field.
func (uuo *UserUpdateOne) SetEmails(s []string) *UserUpdateOne {
	uuo.mutation.SetEmails(s)
	return uuo
}

// AppendEmails appends s to the "emails" field.
func (uuo *UserUpdateOne) AppendEmails(s []string) *UserUpdateOne {
	uuo.mutation.AppendEmails(s)
	return uuo
}

// ClearEmails clears the value of the "emails" field.
func (uuo *UserUpdateOne) ClearEmails() *UserUpdateOne {
	uuo.mutation.ClearEmails()
	return uuo
}

// SetForward sets the "forward" field.
func (uuo *UserUpdateOne) SetForward(b bool) *UserUpdateOne {
	uuo.mutation.SetForward(b)
	return uuo
}

// SetNillableForward sets the "forward" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableForward(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetForward(*b)
	}
	return uuo
}

// SetPaid sets the "paid" field.
func (uuo *UserUpdateOne) SetPaid(b bool) *UserUpdateOne {
	uuo.mutation.SetPaid(b)
	return uuo
}

// SetNillablePaid sets the "paid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePaid(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetPaid(*b)
	}
	return uuo
}

// SetCounter sets the "counter" field.
func (uuo *UserUpdateOne) SetCounter(i int8) *UserUpdateOne {
	uuo.mutation.ResetCounter()
	uuo.mutation.SetCounter(i)
	return uuo
}

// SetNillableCounter sets the "counter" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCounter(i *int8) *UserUpdateOne {
	if i != nil {
		uuo.SetCounter(*i)
	}
	return uuo
}

// AddCounter adds i to the "counter" field.
func (uuo *UserUpdateOne) AddCounter(i int8) *UserUpdateOne {
	uuo.mutation.AddCounter(i)
	return uuo
}

// AddReceiverIDs adds the "receivers" edge to the Receiver entity by IDs.
func (uuo *UserUpdateOne) AddReceiverIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.AddReceiverIDs(ids...)
	return uuo
}

// AddReceivers adds the "receivers" edges to the Receiver entity.
func (uuo *UserUpdateOne) AddReceivers(r ...*Receiver) *UserUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddReceiverIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearReceivers clears all "receivers" edges to the Receiver entity.
func (uuo *UserUpdateOne) ClearReceivers() *UserUpdateOne {
	uuo.mutation.ClearReceivers()
	return uuo
}

// RemoveReceiverIDs removes the "receivers" edge to Receiver entities by IDs.
func (uuo *UserUpdateOne) RemoveReceiverIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.RemoveReceiverIDs(ids...)
	return uuo
}

// RemoveReceivers removes "receivers" edges to Receiver entities.
func (uuo *UserUpdateOne) RemoveReceivers(r ...*Receiver) *UserUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveReceiverIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Emails(); ok {
		_spec.SetField(user.FieldEmails, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedEmails(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldEmails, value)
		})
	}
	if uuo.mutation.EmailsCleared() {
		_spec.ClearField(user.FieldEmails, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Forward(); ok {
		_spec.SetField(user.FieldForward, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Paid(); ok {
		_spec.SetField(user.FieldPaid, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Counter(); ok {
		_spec.SetField(user.FieldCounter, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedCounter(); ok {
		_spec.AddField(user.FieldCounter, field.TypeInt8, value)
	}
	if uuo.mutation.CreatedAtCleared() {
		_spec.ClearField(user.FieldCreatedAt, field.TypeTime)
	}
	if uuo.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !uuo.mutation.ReceiversCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ReceiversIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
