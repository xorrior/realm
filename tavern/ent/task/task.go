// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the lastmodifiedat field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldClaimedAt holds the string denoting the claimedat field in the database.
	FieldClaimedAt = "claimed_at"
	// FieldExecStartedAt holds the string denoting the execstartedat field in the database.
	FieldExecStartedAt = "exec_started_at"
	// FieldExecFinishedAt holds the string denoting the execfinishedat field in the database.
	FieldExecFinishedAt = "exec_finished_at"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// EdgeJob holds the string denoting the job edge name in mutations.
	EdgeJob = "job"
	// EdgeTarget holds the string denoting the target edge name in mutations.
	EdgeTarget = "target"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// JobTable is the table that holds the job relation/edge.
	JobTable = "tasks"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "job_tasks"
	// TargetTable is the table that holds the target relation/edge.
	TargetTable = "tasks"
	// TargetInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetInverseTable = "targets"
	// TargetColumn is the table column denoting the target relation/edge.
	TargetColumn = "task_target"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLastModifiedAt,
	FieldClaimedAt,
	FieldExecStartedAt,
	FieldExecFinishedAt,
	FieldOutput,
	FieldError,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"job_tasks",
	"task_target",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "lastModifiedAt" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "lastModifiedAt" field.
	UpdateDefaultLastModifiedAt func() time.Time
)