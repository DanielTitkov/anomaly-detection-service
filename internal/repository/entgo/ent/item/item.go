// Code generated by entc, DO NOT EDIT.

package item

import (
	"time"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"

	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"

	// Table holds the table name of the item in the database.
	Table = "items"
	// TaskTable is the table the holds the task relation/edge.
	TaskTable = "items"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_items"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSource,
	FieldHash,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Item type.
var ForeignKeys = []string{
	"task_items",
}

var (
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
	// SourceValidator is a validator for the "source" field. It is called by the builders before save.
	SourceValidator func(string) error
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func(string) error
)
