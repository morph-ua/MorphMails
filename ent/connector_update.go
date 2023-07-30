// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"morph_mails/ent/connector"
	"morph_mails/ent/predicate"
	"morph_mails/ent/receiver"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConnectorUpdate is the builder for updating Connector entities.
type ConnectorUpdate struct {
	config
	hooks    []Hook
	mutation *ConnectorMutation
}

// Where appends a list predicates to the ConnectorUpdate builder.
func (cu *ConnectorUpdate) Where(ps ...predicate.Connector) *ConnectorUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ConnectorUpdate) SetName(s string) *ConnectorUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetURL sets the "url" field.
func (cu *ConnectorUpdate) SetURL(s string) *ConnectorUpdate {
	cu.mutation.SetURL(s)
	return cu
}

// SetSecret sets the "secret" field.
func (cu *ConnectorUpdate) SetSecret(s string) *ConnectorUpdate {
	cu.mutation.SetSecret(s)
	return cu
}

// AddReceiverIDs adds the "receivers" edge to the Receiver entity by IDs.
func (cu *ConnectorUpdate) AddReceiverIDs(ids ...string) *ConnectorUpdate {
	cu.mutation.AddReceiverIDs(ids...)
	return cu
}

// AddReceivers adds the "receivers" edges to the Receiver entity.
func (cu *ConnectorUpdate) AddReceivers(r ...*Receiver) *ConnectorUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddReceiverIDs(ids...)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cu *ConnectorUpdate) Mutation() *ConnectorMutation {
	return cu.mutation
}

// ClearReceivers clears all "receivers" edges to the Receiver entity.
func (cu *ConnectorUpdate) ClearReceivers() *ConnectorUpdate {
	cu.mutation.ClearReceivers()
	return cu
}

// RemoveReceiverIDs removes the "receivers" edge to Receiver entities by IDs.
func (cu *ConnectorUpdate) RemoveReceiverIDs(ids ...string) *ConnectorUpdate {
	cu.mutation.RemoveReceiverIDs(ids...)
	return cu
}

// RemoveReceivers removes "receivers" edges to Receiver entities.
func (cu *ConnectorUpdate) RemoveReceivers(r ...*Receiver) *ConnectorUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveReceiverIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConnectorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConnectorUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConnectorUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConnectorUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConnectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(connector.Table, connector.Columns, sqlgraph.NewFieldSpec(connector.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(connector.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.URL(); ok {
		_spec.SetField(connector.FieldURL, field.TypeString, value)
	}
	if value, ok := cu.mutation.Secret(); ok {
		_spec.SetField(connector.FieldSecret, field.TypeString, value)
	}
	if cu.mutation.CreatedAtCleared() {
		_spec.ClearField(connector.FieldCreatedAt, field.TypeTime)
	}
	if cu.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ReceiversTable,
			Columns: []string{connector.ReceiversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receiver.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !cu.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ReceiversTable,
			Columns: []string{connector.ReceiversColumn},
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
	if nodes := cu.mutation.ReceiversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ReceiversTable,
			Columns: []string{connector.ReceiversColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConnectorUpdateOne is the builder for updating a single Connector entity.
type ConnectorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConnectorMutation
}

// SetName sets the "name" field.
func (cuo *ConnectorUpdateOne) SetName(s string) *ConnectorUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetURL sets the "url" field.
func (cuo *ConnectorUpdateOne) SetURL(s string) *ConnectorUpdateOne {
	cuo.mutation.SetURL(s)
	return cuo
}

// SetSecret sets the "secret" field.
func (cuo *ConnectorUpdateOne) SetSecret(s string) *ConnectorUpdateOne {
	cuo.mutation.SetSecret(s)
	return cuo
}

// AddReceiverIDs adds the "receivers" edge to the Receiver entity by IDs.
func (cuo *ConnectorUpdateOne) AddReceiverIDs(ids ...string) *ConnectorUpdateOne {
	cuo.mutation.AddReceiverIDs(ids...)
	return cuo
}

// AddReceivers adds the "receivers" edges to the Receiver entity.
func (cuo *ConnectorUpdateOne) AddReceivers(r ...*Receiver) *ConnectorUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddReceiverIDs(ids...)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cuo *ConnectorUpdateOne) Mutation() *ConnectorMutation {
	return cuo.mutation
}

// ClearReceivers clears all "receivers" edges to the Receiver entity.
func (cuo *ConnectorUpdateOne) ClearReceivers() *ConnectorUpdateOne {
	cuo.mutation.ClearReceivers()
	return cuo
}

// RemoveReceiverIDs removes the "receivers" edge to Receiver entities by IDs.
func (cuo *ConnectorUpdateOne) RemoveReceiverIDs(ids ...string) *ConnectorUpdateOne {
	cuo.mutation.RemoveReceiverIDs(ids...)
	return cuo
}

// RemoveReceivers removes "receivers" edges to Receiver entities.
func (cuo *ConnectorUpdateOne) RemoveReceivers(r ...*Receiver) *ConnectorUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveReceiverIDs(ids...)
}

// Where appends a list predicates to the ConnectorUpdate builder.
func (cuo *ConnectorUpdateOne) Where(ps ...predicate.Connector) *ConnectorUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConnectorUpdateOne) Select(field string, fields ...string) *ConnectorUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Connector entity.
func (cuo *ConnectorUpdateOne) Save(ctx context.Context) (*Connector, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) SaveX(ctx context.Context) *Connector {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConnectorUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConnectorUpdateOne) sqlSave(ctx context.Context) (_node *Connector, err error) {
	_spec := sqlgraph.NewUpdateSpec(connector.Table, connector.Columns, sqlgraph.NewFieldSpec(connector.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Connector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connector.FieldID)
		for _, f := range fields {
			if !connector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != connector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(connector.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.URL(); ok {
		_spec.SetField(connector.FieldURL, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Secret(); ok {
		_spec.SetField(connector.FieldSecret, field.TypeString, value)
	}
	if cuo.mutation.CreatedAtCleared() {
		_spec.ClearField(connector.FieldCreatedAt, field.TypeTime)
	}
	if cuo.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ReceiversTable,
			Columns: []string{connector.ReceiversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receiver.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedReceiversIDs(); len(nodes) > 0 && !cuo.mutation.ReceiversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ReceiversTable,
			Columns: []string{connector.ReceiversColumn},
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
	if nodes := cuo.mutation.ReceiversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   connector.ReceiversTable,
			Columns: []string{connector.ReceiversColumn},
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
	_node = &Connector{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
