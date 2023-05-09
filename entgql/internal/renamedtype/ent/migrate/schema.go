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

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClashingTextsColumns holds the columns for the "clashing_texts" table.
	ClashingTextsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString},
	}
	// ClashingTextsTable holds the schema information for the "clashing_texts" table.
	ClashingTextsTable = &schema.Table{
		Name:       "clashing_texts",
		Columns:    ClashingTextsColumns,
		PrimaryKey: []*schema.Column{ClashingTextsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClashingTextsTable,
	}
)

func init() {
}
