// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"task_manager/ent/hospital"
	"task_manager/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HospitalQuery is the builder for querying Hospital entities.
type HospitalQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Hospital
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HospitalQuery builder.
func (hq *HospitalQuery) Where(ps ...predicate.Hospital) *HospitalQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HospitalQuery) Limit(limit int) *HospitalQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HospitalQuery) Offset(offset int) *HospitalQuery {
	hq.offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HospitalQuery) Unique(unique bool) *HospitalQuery {
	hq.unique = &unique
	return hq
}

// Order adds an order step to the query.
func (hq *HospitalQuery) Order(o ...OrderFunc) *HospitalQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// First returns the first Hospital entity from the query.
// Returns a *NotFoundError when no Hospital was found.
func (hq *HospitalQuery) First(ctx context.Context) (*Hospital, error) {
	nodes, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hospital.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HospitalQuery) FirstX(ctx context.Context) *Hospital {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Hospital ID from the query.
// Returns a *NotFoundError when no Hospital ID was found.
func (hq *HospitalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hospital.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HospitalQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Hospital entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Hospital entity is found.
// Returns a *NotFoundError when no Hospital entities are found.
func (hq *HospitalQuery) Only(ctx context.Context) (*Hospital, error) {
	nodes, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hospital.Label}
	default:
		return nil, &NotSingularError{hospital.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HospitalQuery) OnlyX(ctx context.Context) *Hospital {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Hospital ID in the query.
// Returns a *NotSingularError when more than one Hospital ID is found.
// Returns a *NotFoundError when no entities are found.
func (hq *HospitalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hospital.Label}
	default:
		err = &NotSingularError{hospital.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HospitalQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Hospitals.
func (hq *HospitalQuery) All(ctx context.Context) ([]*Hospital, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HospitalQuery) AllX(ctx context.Context) []*Hospital {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Hospital IDs.
func (hq *HospitalQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hq.Select(hospital.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HospitalQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HospitalQuery) Count(ctx context.Context) (int, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HospitalQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HospitalQuery) Exist(ctx context.Context) (bool, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HospitalQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HospitalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HospitalQuery) Clone() *HospitalQuery {
	if hq == nil {
		return nil
	}
	return &HospitalQuery{
		config:     hq.config,
		limit:      hq.limit,
		offset:     hq.offset,
		order:      append([]OrderFunc{}, hq.order...),
		predicates: append([]predicate.Hospital{}, hq.predicates...),
		// clone intermediate query.
		sql:    hq.sql.Clone(),
		path:   hq.path,
		unique: hq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (hq *HospitalQuery) GroupBy(field string, fields ...string) *HospitalGroupBy {
	grbuild := &HospitalGroupBy{config: hq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(ctx), nil
	}
	grbuild.label = hospital.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (hq *HospitalQuery) Select(fields ...string) *HospitalSelect {
	hq.fields = append(hq.fields, fields...)
	selbuild := &HospitalSelect{HospitalQuery: hq}
	selbuild.label = hospital.Label
	selbuild.flds, selbuild.scan = &hq.fields, selbuild.Scan
	return selbuild
}

func (hq *HospitalQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hq.fields {
		if !hospital.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HospitalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Hospital, error) {
	var (
		nodes = []*Hospital{}
		_spec = hq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Hospital).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Hospital{config: hq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (hq *HospitalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	_spec.Node.Columns = hq.fields
	if len(hq.fields) > 0 {
		_spec.Unique = hq.unique != nil && *hq.unique
	}
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HospitalQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (hq *HospitalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hospital.Table,
			Columns: hospital.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hospital.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if unique := hq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := hq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hospital.FieldID)
		for i := range fields {
			if fields[i] != hospital.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HospitalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(hospital.Table)
	columns := hq.fields
	if len(columns) == 0 {
		columns = hospital.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hq.unique != nil && *hq.unique {
		selector.Distinct()
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HospitalGroupBy is the group-by builder for Hospital entities.
type HospitalGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HospitalGroupBy) Aggregate(fns ...AggregateFunc) *HospitalGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hgb *HospitalGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hgb.path(ctx)
	if err != nil {
		return err
	}
	hgb.sql = query
	return hgb.sqlScan(ctx, v)
}

func (hgb *HospitalGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hgb.fields {
		if !hospital.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HospitalGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql.Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
		for _, f := range hgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(hgb.fields...)...)
}

// HospitalSelect is the builder for selecting fields of Hospital entities.
type HospitalSelect struct {
	*HospitalQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HospitalSelect) Scan(ctx context.Context, v interface{}) error {
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	hs.sql = hs.HospitalQuery.sqlQuery(ctx)
	return hs.sqlScan(ctx, v)
}

func (hs *HospitalSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sql.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}