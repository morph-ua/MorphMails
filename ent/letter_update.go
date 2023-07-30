// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"morph_mails/ent/letter"
	"morph_mails/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LetterUpdate is the builder for updating Letter entities.
type LetterUpdate struct {
	config
	hooks    []Hook
	mutation *LetterMutation
}

// Where appends a list predicates to the LetterUpdate builder.
func (lu *LetterUpdate) Where(ps ...predicate.Letter) *LetterUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetHTML sets the "html" field.
func (lu *LetterUpdate) SetHTML(s string) *LetterUpdate {
	lu.mutation.SetHTML(s)
	return lu
}

// SetFrom sets the "from" field.
func (lu *LetterUpdate) SetFrom(s string) *LetterUpdate {
	lu.mutation.SetFrom(s)
	return lu
}

// SetTo sets the "to" field.
func (lu *LetterUpdate) SetTo(s string) *LetterUpdate {
	lu.mutation.SetTo(s)
	return lu
}

// Mutation returns the LetterMutation object of the builder.
func (lu *LetterUpdate) Mutation() *LetterMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LetterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LetterUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LetterUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LetterUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LetterUpdate) check() error {
	if v, ok := lu.mutation.From(); ok {
		if err := letter.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "Letter.from": %w`, err)}
		}
	}
	if v, ok := lu.mutation.To(); ok {
		if err := letter.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "Letter.to": %w`, err)}
		}
	}
	return nil
}

func (lu *LetterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(letter.Table, letter.Columns, sqlgraph.NewFieldSpec(letter.FieldID, field.TypeString))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.HTML(); ok {
		_spec.SetField(letter.FieldHTML, field.TypeString, value)
	}
	if value, ok := lu.mutation.From(); ok {
		_spec.SetField(letter.FieldFrom, field.TypeString, value)
	}
	if value, ok := lu.mutation.To(); ok {
		_spec.SetField(letter.FieldTo, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{letter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LetterUpdateOne is the builder for updating a single Letter entity.
type LetterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LetterMutation
}

// SetHTML sets the "html" field.
func (luo *LetterUpdateOne) SetHTML(s string) *LetterUpdateOne {
	luo.mutation.SetHTML(s)
	return luo
}

// SetFrom sets the "from" field.
func (luo *LetterUpdateOne) SetFrom(s string) *LetterUpdateOne {
	luo.mutation.SetFrom(s)
	return luo
}

// SetTo sets the "to" field.
func (luo *LetterUpdateOne) SetTo(s string) *LetterUpdateOne {
	luo.mutation.SetTo(s)
	return luo
}

// Mutation returns the LetterMutation object of the builder.
func (luo *LetterUpdateOne) Mutation() *LetterMutation {
	return luo.mutation
}

// Where appends a list predicates to the LetterUpdate builder.
func (luo *LetterUpdateOne) Where(ps ...predicate.Letter) *LetterUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LetterUpdateOne) Select(field string, fields ...string) *LetterUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Letter entity.
func (luo *LetterUpdateOne) Save(ctx context.Context) (*Letter, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LetterUpdateOne) SaveX(ctx context.Context) *Letter {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LetterUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LetterUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LetterUpdateOne) check() error {
	if v, ok := luo.mutation.From(); ok {
		if err := letter.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "Letter.from": %w`, err)}
		}
	}
	if v, ok := luo.mutation.To(); ok {
		if err := letter.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "Letter.to": %w`, err)}
		}
	}
	return nil
}

func (luo *LetterUpdateOne) sqlSave(ctx context.Context) (_node *Letter, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(letter.Table, letter.Columns, sqlgraph.NewFieldSpec(letter.FieldID, field.TypeString))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Letter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, letter.FieldID)
		for _, f := range fields {
			if !letter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != letter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.HTML(); ok {
		_spec.SetField(letter.FieldHTML, field.TypeString, value)
	}
	if value, ok := luo.mutation.From(); ok {
		_spec.SetField(letter.FieldFrom, field.TypeString, value)
	}
	if value, ok := luo.mutation.To(); ok {
		_spec.SetField(letter.FieldTo, field.TypeString, value)
	}
	_node = &Letter{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{letter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
