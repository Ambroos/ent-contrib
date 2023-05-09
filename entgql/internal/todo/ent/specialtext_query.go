// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entgql/internal/todo/ent/predicate"
	"entgo.io/contrib/entgql/internal/todo/ent/specialtext"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpecialTextQuery is the builder for querying SpecialText entities.
type SpecialTextQuery struct {
	config
	ctx        *QueryContext
	order      []specialtext.OrderOption
	inters     []Interceptor
	predicates []predicate.SpecialText
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*SpecialText) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SpecialTextQuery builder.
func (stq *SpecialTextQuery) Where(ps ...predicate.SpecialText) *SpecialTextQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit the number of records to be returned by this query.
func (stq *SpecialTextQuery) Limit(limit int) *SpecialTextQuery {
	stq.ctx.Limit = &limit
	return stq
}

// Offset to start from.
func (stq *SpecialTextQuery) Offset(offset int) *SpecialTextQuery {
	stq.ctx.Offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *SpecialTextQuery) Unique(unique bool) *SpecialTextQuery {
	stq.ctx.Unique = &unique
	return stq
}

// Order specifies how the records should be ordered.
func (stq *SpecialTextQuery) Order(o ...specialtext.OrderOption) *SpecialTextQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// First returns the first SpecialText entity from the query.
// Returns a *NotFoundError when no SpecialText was found.
func (stq *SpecialTextQuery) First(ctx context.Context) (*SpecialText, error) {
	nodes, err := stq.Limit(1).All(setContextOp(ctx, stq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{specialtext.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *SpecialTextQuery) FirstX(ctx context.Context) *SpecialText {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SpecialText ID from the query.
// Returns a *NotFoundError when no SpecialText ID was found.
func (stq *SpecialTextQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(1).IDs(setContextOp(ctx, stq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{specialtext.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *SpecialTextQuery) FirstIDX(ctx context.Context) int {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SpecialText entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SpecialText entity is found.
// Returns a *NotFoundError when no SpecialText entities are found.
func (stq *SpecialTextQuery) Only(ctx context.Context) (*SpecialText, error) {
	nodes, err := stq.Limit(2).All(setContextOp(ctx, stq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{specialtext.Label}
	default:
		return nil, &NotSingularError{specialtext.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *SpecialTextQuery) OnlyX(ctx context.Context) *SpecialText {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SpecialText ID in the query.
// Returns a *NotSingularError when more than one SpecialText ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *SpecialTextQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(2).IDs(setContextOp(ctx, stq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{specialtext.Label}
	default:
		err = &NotSingularError{specialtext.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *SpecialTextQuery) OnlyIDX(ctx context.Context) int {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SpecialTexts.
func (stq *SpecialTextQuery) All(ctx context.Context) ([]*SpecialText, error) {
	ctx = setContextOp(ctx, stq.ctx, "All")
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SpecialText, *SpecialTextQuery]()
	return withInterceptors[[]*SpecialText](ctx, stq, qr, stq.inters)
}

// AllX is like All, but panics if an error occurs.
func (stq *SpecialTextQuery) AllX(ctx context.Context) []*SpecialText {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SpecialText IDs.
func (stq *SpecialTextQuery) IDs(ctx context.Context) (ids []int, err error) {
	if stq.ctx.Unique == nil && stq.path != nil {
		stq.Unique(true)
	}
	ctx = setContextOp(ctx, stq.ctx, "IDs")
	if err = stq.Select(specialtext.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *SpecialTextQuery) IDsX(ctx context.Context) []int {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *SpecialTextQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, stq.ctx, "Count")
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, stq, querierCount[*SpecialTextQuery](), stq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (stq *SpecialTextQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *SpecialTextQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, stq.ctx, "Exist")
	switch _, err := stq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *SpecialTextQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpecialTextQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *SpecialTextQuery) Clone() *SpecialTextQuery {
	if stq == nil {
		return nil
	}
	return &SpecialTextQuery{
		config:     stq.config,
		ctx:        stq.ctx.Clone(),
		order:      append([]specialtext.OrderOption{}, stq.order...),
		inters:     append([]Interceptor{}, stq.inters...),
		predicates: append([]predicate.SpecialText{}, stq.predicates...),
		// clone intermediate query.
		sql:  stq.sql.Clone(),
		path: stq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SpecialText.Query().
//		GroupBy(specialtext.FieldContent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (stq *SpecialTextQuery) GroupBy(field string, fields ...string) *SpecialTextGroupBy {
	stq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SpecialTextGroupBy{build: stq}
	grbuild.flds = &stq.ctx.Fields
	grbuild.label = specialtext.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//	}
//
//	client.SpecialText.Query().
//		Select(specialtext.FieldContent).
//		Scan(ctx, &v)
func (stq *SpecialTextQuery) Select(fields ...string) *SpecialTextSelect {
	stq.ctx.Fields = append(stq.ctx.Fields, fields...)
	sbuild := &SpecialTextSelect{SpecialTextQuery: stq}
	sbuild.label = specialtext.Label
	sbuild.flds, sbuild.scan = &stq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SpecialTextSelect configured with the given aggregations.
func (stq *SpecialTextQuery) Aggregate(fns ...AggregateFunc) *SpecialTextSelect {
	return stq.Select().Aggregate(fns...)
}

func (stq *SpecialTextQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range stq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, stq); err != nil {
				return err
			}
		}
	}
	for _, f := range stq.ctx.Fields {
		if !specialtext.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *SpecialTextQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SpecialText, error) {
	var (
		nodes = []*SpecialText{}
		_spec = stq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SpecialText).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SpecialText{config: stq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(stq.modifiers) > 0 {
		_spec.Modifiers = stq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range stq.loadTotal {
		if err := stq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (stq *SpecialTextQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	if len(stq.modifiers) > 0 {
		_spec.Modifiers = stq.modifiers
	}
	_spec.Node.Columns = stq.ctx.Fields
	if len(stq.ctx.Fields) > 0 {
		_spec.Unique = stq.ctx.Unique != nil && *stq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *SpecialTextQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(specialtext.Table, specialtext.Columns, sqlgraph.NewFieldSpec(specialtext.FieldID, field.TypeInt))
	_spec.From = stq.sql
	if unique := stq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if stq.path != nil {
		_spec.Unique = true
	}
	if fields := stq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, specialtext.FieldID)
		for i := range fields {
			if fields[i] != specialtext.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *SpecialTextQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(specialtext.Table)
	columns := stq.ctx.Fields
	if len(columns) == 0 {
		columns = specialtext.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.ctx.Unique != nil && *stq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SpecialTextGroupBy is the group-by builder for SpecialText entities.
type SpecialTextGroupBy struct {
	selector
	build *SpecialTextQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *SpecialTextGroupBy) Aggregate(fns ...AggregateFunc) *SpecialTextGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the selector query and scans the result into the given value.
func (stgb *SpecialTextGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, stgb.build.ctx, "GroupBy")
	if err := stgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecialTextQuery, *SpecialTextGroupBy](ctx, stgb.build, stgb, stgb.build.inters, v)
}

func (stgb *SpecialTextGroupBy) sqlScan(ctx context.Context, root *SpecialTextQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*stgb.flds)+len(stgb.fns))
		for _, f := range *stgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*stgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SpecialTextSelect is the builder for selecting fields of SpecialText entities.
type SpecialTextSelect struct {
	*SpecialTextQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sts *SpecialTextSelect) Aggregate(fns ...AggregateFunc) *SpecialTextSelect {
	sts.fns = append(sts.fns, fns...)
	return sts
}

// Scan applies the selector query and scans the result into the given value.
func (sts *SpecialTextSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sts.ctx, "Select")
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpecialTextQuery, *SpecialTextSelect](ctx, sts.SpecialTextQuery, sts, sts.inters, v)
}

func (sts *SpecialTextSelect) sqlScan(ctx context.Context, root *SpecialTextQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sts.fns))
	for _, fn := range sts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
