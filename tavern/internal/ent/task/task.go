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
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldClaimedAt holds the string denoting the claimed_at field in the database.
	FieldClaimedAt = "claimed_at"
	// FieldExecStartedAt holds the string denoting the exec_started_at field in the database.
	FieldExecStartedAt = "exec_started_at"
	// FieldExecFinishedAt holds the string denoting the exec_finished_at field in the database.
	FieldExecFinishedAt = "exec_finished_at"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// EdgeQuest holds the string denoting the quest edge name in mutations.
	EdgeQuest = "quest"
	// EdgeBeacon holds the string denoting the beacon edge name in mutations.
	EdgeBeacon = "beacon"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// QuestTable is the table that holds the quest relation/edge.
	QuestTable = "tasks"
	// QuestInverseTable is the table name for the Quest entity.
	// It exists in this package in order to avoid circular dependency with the "quest" package.
	QuestInverseTable = "quests"
	// QuestColumn is the table column denoting the quest relation/edge.
	QuestColumn = "quest_tasks"
	// BeaconTable is the table that holds the beacon relation/edge.
	BeaconTable = "tasks"
	// BeaconInverseTable is the table name for the Beacon entity.
	// It exists in this package in order to avoid circular dependency with the "beacon" package.
	BeaconInverseTable = "beacons"
	// BeaconColumn is the table column denoting the beacon relation/edge.
	BeaconColumn = "task_beacon"
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
	"quest_tasks",
	"task_beacon",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "last_modified_at" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "last_modified_at" field.
	UpdateDefaultLastModifiedAt func() time.Time
)
