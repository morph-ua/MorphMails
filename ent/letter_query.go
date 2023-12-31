// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"morph_mails/ent/letter"
	"morph_mails/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LetterQuery is the builder for querying Letter entities.
type LetterQuery struct {
	config
	ctx        *QueryContext
	order      []letter.OrderOption
	inters     []Interceptor
	predicates []predicate.Letter
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LetterQuery builder.
func (lq *LetterQuery) Where(ps ...predicate.Letter) *LetterQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LetterQuery) Limit(limit int) *LetterQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LetterQuery) Offset(offset int) *LetterQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LetterQuery) Unique(unique bool) *LetterQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LetterQuery) Order(o ...letter.OrderOption) *LetterQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// First returns the first Letter entity from the query.
// Returns a *NotFoundError when no Letter was found.
func (lq *LetterQuery) First(ctx context.Context) (*Letter, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{letter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LetterQuery) FirstX(ctx context.Context) *Letter {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Letter ID from the query.
// Returns a *NotFoundError when no Letter ID was found.
func (lq *LetterQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{letter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LetterQuery) FirstIDX(ctx context.Context) string {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Letter entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Letter entity is found.
// Returns a *NotFoundError when no Letter entities are found.
func (lq *LetterQuery) Only(ctx context.Context) (*Letter, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{letter.Label}
	default:
		return nil, &NotSingularError{letter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LetterQuery) OnlyX(ctx context.Context) *Letter {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Letter ID in the query.
// Returns a *NotSingularError when more than one Letter ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LetterQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{letter.Label}
	default:
		err = &NotSingularError{letter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LetterQuery) OnlyIDX(ctx context.Context) string {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Letters.
func (lq *LetterQuery) All(ctx context.Context) ([]*Letter, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Letter, *LetterQuery]()
	return withInterceptors[[]*Letter](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LetterQuery) AllX(ctx context.Context) []*Letter {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Letter IDs.
func (lq *LetterQuery) IDs(ctx context.Context) (ids []string, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(letter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LetterQuery) IDsX(ctx context.Context) []string {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LetterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LetterQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LetterQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LetterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LetterQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LetterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LetterQuery) Clone() *LetterQuery {
	if lq == nil {
		return nil
	}
	return &LetterQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]letter.OrderOption{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.Letter{}, lq.predicates...),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HTML string `json:"html,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Letter.Query().
//		GroupBy(letter.FieldHTML).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LetterQuery) GroupBy(field string, fields ...string) *LetterGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LetterGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = letter.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HTML string `json:"html,omitempty"`
//	}
//
//	client.Letter.Query().
//		Select(letter.FieldHTML).
//		Scan(ctx, &v)
func (lq *LetterQuery) Select(fields ...string) *LetterSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LetterSelect{LetterQuery: lq}
	sbuild.label = letter.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LetterSelect configured with the given aggregations.
func (lq *LetterQuery) Aggregate(fns ...AggregateFunc) *LetterSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LetterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !letter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LetterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Letter, error) {
	var (
		nodes = []*Letter{}
		_spec = lq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Letter).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Letter{config: lq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lq *LetterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LetterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(letter.Table, letter.Columns, sqlgraph.NewFieldSpec(letter.FieldID, field.TypeString))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, letter.FieldID)
		for i := range fields {
			if fields[i] != letter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LetterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(letter.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = letter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LetterGroupBy is the group-by builder for Letter entities.
type LetterGroupBy struct {
	selector
	build *LetterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LetterGroupBy) Aggregate(fns ...AggregateFunc) *LetterGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LetterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LetterQuery, *LetterGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LetterGroupBy) sqlScan(ctx context.Context, root *LetterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LetterSelect is the builder for selecting fields of Letter entities.
type LetterSelect struct {
	*LetterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LetterSelect) Aggregate(fns ...AggregateFunc) *LetterSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LetterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LetterQuery, *LetterSelect](ctx, ls.LetterQuery, ls, ls.inters, v)
}

func (ls *LetterSelect) sqlScan(ctx context.Context, root *LetterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
