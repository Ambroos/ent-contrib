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
	"errors"
	"fmt"

	"entgo.io/contrib/entgql/internal/renamedtype/ent/clashingtext"
	"entgo.io/contrib/entgql/internal/renamedtype/ent/predicate"
)

// NotClashingTextWhereInput represents a where input for filtering ClashingText queries.
type NotClashingTextWhereInput struct {
	Predicates []predicate.ClashingText     `json:"-"`
	Not        *NotClashingTextWhereInput   `json:"not,omitempty"`
	Or         []*NotClashingTextWhereInput `json:"or,omitempty"`
	And        []*NotClashingTextWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "content" field predicates.
	Content             *string  `json:"content,omitempty"`
	ContentNEQ          *string  `json:"contentNEQ,omitempty"`
	ContentIn           []string `json:"contentIn,omitempty"`
	ContentNotIn        []string `json:"contentNotIn,omitempty"`
	ContentGT           *string  `json:"contentGT,omitempty"`
	ContentGTE          *string  `json:"contentGTE,omitempty"`
	ContentLT           *string  `json:"contentLT,omitempty"`
	ContentLTE          *string  `json:"contentLTE,omitempty"`
	ContentContains     *string  `json:"contentContains,omitempty"`
	ContentHasPrefix    *string  `json:"contentHasPrefix,omitempty"`
	ContentHasSuffix    *string  `json:"contentHasSuffix,omitempty"`
	ContentEqualFold    *string  `json:"contentEqualFold,omitempty"`
	ContentContainsFold *string  `json:"contentContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *NotClashingTextWhereInput) AddPredicates(predicates ...predicate.ClashingText) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the NotClashingTextWhereInput filter on the ClashingTextQuery builder.
func (i *NotClashingTextWhereInput) Filter(q *ClashingTextQuery) (*ClashingTextQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyNotClashingTextWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyNotClashingTextWhereInput is returned in case the NotClashingTextWhereInput is empty.
var ErrEmptyNotClashingTextWhereInput = errors.New("ent: empty predicate NotClashingTextWhereInput")

// P returns a predicate for filtering clashingtexts.
// An error is returned if the input is empty or invalid.
func (i *NotClashingTextWhereInput) P() (predicate.ClashingText, error) {
	var predicates []predicate.ClashingText
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, clashingtext.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.ClashingText, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, clashingtext.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.ClashingText, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, clashingtext.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, clashingtext.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, clashingtext.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, clashingtext.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, clashingtext.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, clashingtext.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, clashingtext.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, clashingtext.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, clashingtext.IDLTE(*i.IDLTE))
	}
	if i.Content != nil {
		predicates = append(predicates, clashingtext.ContentEQ(*i.Content))
	}
	if i.ContentNEQ != nil {
		predicates = append(predicates, clashingtext.ContentNEQ(*i.ContentNEQ))
	}
	if len(i.ContentIn) > 0 {
		predicates = append(predicates, clashingtext.ContentIn(i.ContentIn...))
	}
	if len(i.ContentNotIn) > 0 {
		predicates = append(predicates, clashingtext.ContentNotIn(i.ContentNotIn...))
	}
	if i.ContentGT != nil {
		predicates = append(predicates, clashingtext.ContentGT(*i.ContentGT))
	}
	if i.ContentGTE != nil {
		predicates = append(predicates, clashingtext.ContentGTE(*i.ContentGTE))
	}
	if i.ContentLT != nil {
		predicates = append(predicates, clashingtext.ContentLT(*i.ContentLT))
	}
	if i.ContentLTE != nil {
		predicates = append(predicates, clashingtext.ContentLTE(*i.ContentLTE))
	}
	if i.ContentContains != nil {
		predicates = append(predicates, clashingtext.ContentContains(*i.ContentContains))
	}
	if i.ContentHasPrefix != nil {
		predicates = append(predicates, clashingtext.ContentHasPrefix(*i.ContentHasPrefix))
	}
	if i.ContentHasSuffix != nil {
		predicates = append(predicates, clashingtext.ContentHasSuffix(*i.ContentHasSuffix))
	}
	if i.ContentEqualFold != nil {
		predicates = append(predicates, clashingtext.ContentEqualFold(*i.ContentEqualFold))
	}
	if i.ContentContainsFold != nil {
		predicates = append(predicates, clashingtext.ContentContainsFold(*i.ContentContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyNotClashingTextWhereInput
	case 1:
		return predicates[0], nil
	default:
		return clashingtext.And(predicates...), nil
	}
}
