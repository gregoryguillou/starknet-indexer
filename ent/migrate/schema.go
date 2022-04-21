// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlocksColumns holds the columns for the "blocks" table.
	BlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "block_hash", Type: field.TypeString, Unique: true},
		{Name: "parent_block_hash", Type: field.TypeString},
		{Name: "block_number", Type: field.TypeUint64, Unique: true},
		{Name: "state_root", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// BlocksTable holds the schema information for the "blocks" table.
	BlocksTable = &schema.Table{
		Name:       "blocks",
		Columns:    BlocksColumns,
		PrimaryKey: []*schema.Column{BlocksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlocksTable,
	}
)

func init() {
}
