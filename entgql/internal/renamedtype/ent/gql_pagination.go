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
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entgql/internal/renamedtype/ent/clashingtext"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// NotClashingText is the type alias for ClashingText.
type NotClashingText = ClashingText

// NotClashingTextEdge is the edge representation of NotClashingText.
type NotClashingTextEdge struct {
	Node   *NotClashingText `json:"node"`
	Cursor Cursor           `json:"cursor"`
}

// NotClashingTextConnection is the connection containing edges to NotClashingText.
type NotClashingTextConnection struct {
	Edges      []*NotClashingTextEdge `json:"edges"`
	PageInfo   PageInfo               `json:"pageInfo"`
	TotalCount int                    `json:"totalCount"`
}

func (c *NotClashingTextConnection) build(nodes []*NotClashingText, pager *notclashingtextPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *NotClashingText
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *NotClashingText {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *NotClashingText {
			return nodes[i]
		}
	}
	c.Edges = make([]*NotClashingTextEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &NotClashingTextEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// NotClashingTextPaginateOption enables pagination customization.
type NotClashingTextPaginateOption func(*notclashingtextPager) error

// WithNotClashingTextOrder configures pagination ordering.
func WithNotClashingTextOrder(order *NotClashingTextOrder) NotClashingTextPaginateOption {
	if order == nil {
		order = DefaultNotClashingTextOrder
	}
	o := *order
	return func(pager *notclashingtextPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultNotClashingTextOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithNotClashingTextFilter configures pagination filter.
func WithNotClashingTextFilter(filter func(*ClashingTextQuery) (*ClashingTextQuery, error)) NotClashingTextPaginateOption {
	return func(pager *notclashingtextPager) error {
		if filter == nil {
			return errors.New("ClashingTextQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type notclashingtextPager struct {
	reverse bool
	order   *NotClashingTextOrder
	filter  func(*ClashingTextQuery) (*ClashingTextQuery, error)
}

func newNotClashingTextPager(opts []NotClashingTextPaginateOption, reverse bool) (*notclashingtextPager, error) {
	pager := &notclashingtextPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultNotClashingTextOrder
	}
	return pager, nil
}

func (p *notclashingtextPager) applyFilter(query *ClashingTextQuery) (*ClashingTextQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *notclashingtextPager) toCursor(ct *NotClashingText) Cursor {
	return p.order.Field.toCursor(ct)
}

func (p *notclashingtextPager) applyCursors(query *ClashingTextQuery, after, before *Cursor) (*ClashingTextQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultNotClashingTextOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *notclashingtextPager) applyOrder(query *ClashingTextQuery) *ClashingTextQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultNotClashingTextOrder.Field {
		query = query.Order(DefaultNotClashingTextOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *notclashingtextPager) orderExpr(query *ClashingTextQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultNotClashingTextOrder.Field {
			b.Comma().Ident(DefaultNotClashingTextOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to NotClashingText.
func (ct *ClashingTextQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...NotClashingTextPaginateOption,
) (*NotClashingTextConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newNotClashingTextPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if ct, err = pager.applyFilter(ct); err != nil {
		return nil, err
	}
	conn := &NotClashingTextConnection{Edges: []*NotClashingTextEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = ct.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if ct, err = pager.applyCursors(ct, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		ct.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ct.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	ct = pager.applyOrder(ct)
	nodes, err := ct.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// NotClashingTextOrderField defines the ordering field of ClashingText.
type NotClashingTextOrderField struct {
	// Value extracts the ordering value from the given ClashingText.
	Value    func(*NotClashingText) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) clashingtext.OrderOption
	toCursor func(*NotClashingText) Cursor
}

// NotClashingTextOrder defines the ordering of ClashingText.
type NotClashingTextOrder struct {
	Direction OrderDirection             `json:"direction"`
	Field     *NotClashingTextOrderField `json:"field"`
}

// DefaultNotClashingTextOrder is the default ordering of ClashingText.
var DefaultNotClashingTextOrder = &NotClashingTextOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &NotClashingTextOrderField{
		Value: func(ct *NotClashingText) (ent.Value, error) {
			return ct.ID, nil
		},
		column: clashingtext.FieldID,
		toTerm: clashingtext.ByID,
		toCursor: func(ct *NotClashingText) Cursor {
			return Cursor{ID: ct.ID}
		},
	},
}

// ToEdge converts NotClashingText into NotClashingTextEdge.
func (ct *NotClashingText) ToEdge(order *NotClashingTextOrder) *NotClashingTextEdge {
	if order == nil {
		order = DefaultNotClashingTextOrder
	}
	return &NotClashingTextEdge{
		Node:   ct,
		Cursor: order.Field.toCursor(ct),
	}
}
