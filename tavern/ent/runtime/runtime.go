// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/kcarretto/realm/tavern/ent/file"
	"github.com/kcarretto/realm/tavern/ent/job"
	"github.com/kcarretto/realm/tavern/ent/schema"
	"github.com/kcarretto/realm/tavern/ent/session"
	"github.com/kcarretto/realm/tavern/ent/tag"
	"github.com/kcarretto/realm/tavern/ent/task"
	"github.com/kcarretto/realm/tavern/ent/tome"
	"github.com/kcarretto/realm/tavern/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileMixin := schema.File{}.Mixin()
	fileHooks := schema.File{}.Hooks()
	file.Hooks[0] = fileHooks[0]
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileMixinFields0[0].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescLastModifiedAt is the schema descriptor for last_modified_at field.
	fileDescLastModifiedAt := fileMixinFields0[1].Descriptor()
	// file.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	file.DefaultLastModifiedAt = fileDescLastModifiedAt.Default.(func() time.Time)
	// file.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	file.UpdateDefaultLastModifiedAt = fileDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// fileDescName is the schema descriptor for name field.
	fileDescName := fileFields[0].Descriptor()
	// file.NameValidator is a validator for the "name" field. It is called by the builders before save.
	file.NameValidator = fileDescName.Validators[0].(func(string) error)
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[1].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	// fileDescHash is the schema descriptor for hash field.
	fileDescHash := fileFields[2].Descriptor()
	// file.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	file.HashValidator = fileDescHash.Validators[0].(func(string) error)
	jobMixin := schema.Job{}.Mixin()
	jobMixinFields0 := jobMixin[0].Fields()
	_ = jobMixinFields0
	jobFields := schema.Job{}.Fields()
	_ = jobFields
	// jobDescCreatedAt is the schema descriptor for created_at field.
	jobDescCreatedAt := jobMixinFields0[0].Descriptor()
	// job.DefaultCreatedAt holds the default value on creation for the created_at field.
	job.DefaultCreatedAt = jobDescCreatedAt.Default.(func() time.Time)
	// jobDescLastModifiedAt is the schema descriptor for last_modified_at field.
	jobDescLastModifiedAt := jobMixinFields0[1].Descriptor()
	// job.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	job.DefaultLastModifiedAt = jobDescLastModifiedAt.Default.(func() time.Time)
	// job.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	job.UpdateDefaultLastModifiedAt = jobDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// jobDescName is the schema descriptor for name field.
	jobDescName := jobFields[0].Descriptor()
	// job.NameValidator is a validator for the "name" field. It is called by the builders before save.
	job.NameValidator = jobDescName.Validators[0].(func(string) error)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescName is the schema descriptor for name field.
	sessionDescName := sessionFields[0].Descriptor()
	// session.DefaultName holds the default value on creation for the name field.
	session.DefaultName = sessionDescName.Default.(func() string)
	// session.NameValidator is a validator for the "name" field. It is called by the builders before save.
	session.NameValidator = sessionDescName.Validators[0].(func(string) error)
	// sessionDescPrincipal is the schema descriptor for principal field.
	sessionDescPrincipal := sessionFields[1].Descriptor()
	// session.PrincipalValidator is a validator for the "principal" field. It is called by the builders before save.
	session.PrincipalValidator = sessionDescPrincipal.Validators[0].(func(string) error)
	// sessionDescHostname is the schema descriptor for hostname field.
	sessionDescHostname := sessionFields[2].Descriptor()
	// session.HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	session.HostnameValidator = sessionDescHostname.Validators[0].(func(string) error)
	// sessionDescIdentifier is the schema descriptor for identifier field.
	sessionDescIdentifier := sessionFields[3].Descriptor()
	// session.DefaultIdentifier holds the default value on creation for the identifier field.
	session.DefaultIdentifier = sessionDescIdentifier.Default.(func() string)
	// session.IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	session.IdentifierValidator = sessionDescIdentifier.Validators[0].(func(string) error)
	// sessionDescAgentIdentifier is the schema descriptor for agent_identifier field.
	sessionDescAgentIdentifier := sessionFields[4].Descriptor()
	// session.AgentIdentifierValidator is a validator for the "agent_identifier" field. It is called by the builders before save.
	session.AgentIdentifierValidator = sessionDescAgentIdentifier.Validators[0].(func(string) error)
	// sessionDescHostIdentifier is the schema descriptor for host_identifier field.
	sessionDescHostIdentifier := sessionFields[5].Descriptor()
	// session.HostIdentifierValidator is a validator for the "host_identifier" field. It is called by the builders before save.
	session.HostIdentifierValidator = sessionDescHostIdentifier.Validators[0].(func(string) error)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[0].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskMixinFields0[0].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescLastModifiedAt is the schema descriptor for last_modified_at field.
	taskDescLastModifiedAt := taskMixinFields0[1].Descriptor()
	// task.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	task.DefaultLastModifiedAt = taskDescLastModifiedAt.Default.(func() time.Time)
	// task.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	task.UpdateDefaultLastModifiedAt = taskDescLastModifiedAt.UpdateDefault.(func() time.Time)
	tomeMixin := schema.Tome{}.Mixin()
	tomeHooks := schema.Tome{}.Hooks()
	tome.Hooks[0] = tomeHooks[0]
	tomeMixinFields0 := tomeMixin[0].Fields()
	_ = tomeMixinFields0
	tomeFields := schema.Tome{}.Fields()
	_ = tomeFields
	// tomeDescCreatedAt is the schema descriptor for created_at field.
	tomeDescCreatedAt := tomeMixinFields0[0].Descriptor()
	// tome.DefaultCreatedAt holds the default value on creation for the created_at field.
	tome.DefaultCreatedAt = tomeDescCreatedAt.Default.(func() time.Time)
	// tomeDescLastModifiedAt is the schema descriptor for last_modified_at field.
	tomeDescLastModifiedAt := tomeMixinFields0[1].Descriptor()
	// tome.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	tome.DefaultLastModifiedAt = tomeDescLastModifiedAt.Default.(func() time.Time)
	// tome.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	tome.UpdateDefaultLastModifiedAt = tomeDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// tomeDescName is the schema descriptor for name field.
	tomeDescName := tomeFields[0].Descriptor()
	// tome.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tome.NameValidator = tomeDescName.Validators[0].(func(string) error)
	// tomeDescHash is the schema descriptor for hash field.
	tomeDescHash := tomeFields[3].Descriptor()
	// tome.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	tome.HashValidator = tomeDescHash.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescSessionToken is the schema descriptor for session_token field.
	userDescSessionToken := userFields[3].Descriptor()
	// user.DefaultSessionToken holds the default value on creation for the session_token field.
	user.DefaultSessionToken = userDescSessionToken.Default.(func() string)
	// user.SessionTokenValidator is a validator for the "session_token" field. It is called by the builders before save.
	user.SessionTokenValidator = userDescSessionToken.Validators[0].(func(string) error)
	// userDescIsActivated is the schema descriptor for is_activated field.
	userDescIsActivated := userFields[4].Descriptor()
	// user.DefaultIsActivated holds the default value on creation for the is_activated field.
	user.DefaultIsActivated = userDescIsActivated.Default.(bool)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[5].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}

const (
	Version = "v0.11.9"                                         // Version of ent codegen.
	Sum     = "h1:dbbCkAiPVTRBIJwoZctiSYjB7zxQIBOzVSU5H9VYIQI=" // Sum of ent codegen.
)
