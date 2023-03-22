// Code generated by ent, DO NOT EDIT.

package session

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrincipal holds the string denoting the principal field in the database.
	FieldPrincipal = "principal"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldAgentIdentifier holds the string denoting the agent_identifier field in the database.
	FieldAgentIdentifier = "agent_identifier"
	// FieldHostIdentifier holds the string denoting the host_identifier field in the database.
	FieldHostIdentifier = "host_identifier"
	// FieldHostPrimaryIP holds the string denoting the host_primary_ip field in the database.
	FieldHostPrimaryIP = "host_primary_ip"
	// FieldHostPlatform holds the string denoting the host_platform field in the database.
	FieldHostPlatform = "host_platform"
	// FieldLastSeenAt holds the string denoting the last_seen_at field in the database.
	FieldLastSeenAt = "last_seen_at"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// Table holds the table name of the session in the database.
	Table = "sessions"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "session_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "task_session"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPrincipal,
	FieldHostname,
	FieldIdentifier,
	FieldAgentIdentifier,
	FieldHostIdentifier,
	FieldHostPrimaryIP,
	FieldHostPlatform,
	FieldLastSeenAt,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"session_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName func() string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PrincipalValidator is a validator for the "principal" field. It is called by the builders before save.
	PrincipalValidator func(string) error
	// HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	HostnameValidator func(string) error
	// DefaultIdentifier holds the default value on creation for the "identifier" field.
	DefaultIdentifier func() string
	// IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	IdentifierValidator func(string) error
	// AgentIdentifierValidator is a validator for the "agent_identifier" field. It is called by the builders before save.
	AgentIdentifierValidator func(string) error
	// HostIdentifierValidator is a validator for the "host_identifier" field. It is called by the builders before save.
	HostIdentifierValidator func(string) error
)

// HostPlatform defines the type for the "host_platform" enum field.
type HostPlatform string

// HostPlatformUnknown is the default value of the HostPlatform enum.
const DefaultHostPlatform = HostPlatformUnknown

// HostPlatform values.
const (
	HostPlatformWindows HostPlatform = "Windows"
	HostPlatformLinux   HostPlatform = "Linux"
	HostPlatformMacOS   HostPlatform = "MacOS"
	HostPlatformBSD     HostPlatform = "BSD"
	HostPlatformUnknown HostPlatform = "Unknown"
)

func (hp HostPlatform) String() string {
	return string(hp)
}

// HostPlatformValidator is a validator for the "host_platform" field enum values. It is called by the builders before save.
func HostPlatformValidator(hp HostPlatform) error {
	switch hp {
	case HostPlatformWindows, HostPlatformLinux, HostPlatformMacOS, HostPlatformBSD, HostPlatformUnknown:
		return nil
	default:
		return fmt.Errorf("session: invalid enum value for host_platform field: %q", hp)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e HostPlatform) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *HostPlatform) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = HostPlatform(str)
	if err := HostPlatformValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid HostPlatform", str)
	}
	return nil
}
