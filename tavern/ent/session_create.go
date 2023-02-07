// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/session"
	"github.com/kcarretto/realm/tavern/ent/tag"
	"github.com/kcarretto/realm/tavern/ent/task"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetPrincipal sets the "principal" field.
func (sc *SessionCreate) SetPrincipal(s string) *SessionCreate {
	sc.mutation.SetPrincipal(s)
	return sc
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (sc *SessionCreate) SetNillablePrincipal(s *string) *SessionCreate {
	if s != nil {
		sc.SetPrincipal(*s)
	}
	return sc
}

// SetHostname sets the "hostname" field.
func (sc *SessionCreate) SetHostname(s string) *SessionCreate {
	sc.mutation.SetHostname(s)
	return sc
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (sc *SessionCreate) SetNillableHostname(s *string) *SessionCreate {
	if s != nil {
		sc.SetHostname(*s)
	}
	return sc
}

// SetIdentifier sets the "identifier" field.
func (sc *SessionCreate) SetIdentifier(s string) *SessionCreate {
	sc.mutation.SetIdentifier(s)
	return sc
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (sc *SessionCreate) SetNillableIdentifier(s *string) *SessionCreate {
	if s != nil {
		sc.SetIdentifier(*s)
	}
	return sc
}

// SetAgentIdentifier sets the "agentIdentifier" field.
func (sc *SessionCreate) SetAgentIdentifier(s string) *SessionCreate {
	sc.mutation.SetAgentIdentifier(s)
	return sc
}

// SetNillableAgentIdentifier sets the "agentIdentifier" field if the given value is not nil.
func (sc *SessionCreate) SetNillableAgentIdentifier(s *string) *SessionCreate {
	if s != nil {
		sc.SetAgentIdentifier(*s)
	}
	return sc
}

// SetHostIdentifier sets the "hostIdentifier" field.
func (sc *SessionCreate) SetHostIdentifier(s string) *SessionCreate {
	sc.mutation.SetHostIdentifier(s)
	return sc
}

// SetNillableHostIdentifier sets the "hostIdentifier" field if the given value is not nil.
func (sc *SessionCreate) SetNillableHostIdentifier(s *string) *SessionCreate {
	if s != nil {
		sc.SetHostIdentifier(*s)
	}
	return sc
}

// SetLastSeenAt sets the "lastSeenAt" field.
func (sc *SessionCreate) SetLastSeenAt(t time.Time) *SessionCreate {
	sc.mutation.SetLastSeenAt(t)
	return sc
}

// SetNillableLastSeenAt sets the "lastSeenAt" field if the given value is not nil.
func (sc *SessionCreate) SetNillableLastSeenAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetLastSeenAt(*t)
	}
	return sc
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (sc *SessionCreate) AddTagIDs(ids ...int) *SessionCreate {
	sc.mutation.AddTagIDs(ids...)
	return sc
}

// AddTags adds the "tags" edges to the Tag entity.
func (sc *SessionCreate) AddTags(t ...*Tag) *SessionCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTagIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (sc *SessionCreate) AddTaskIDs(ids ...int) *SessionCreate {
	sc.mutation.AddTaskIDs(ids...)
	return sc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (sc *SessionCreate) AddTasks(t ...*Task) *SessionCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTaskIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.Identifier(); !ok {
		v := session.DefaultIdentifier()
		sc.mutation.SetIdentifier(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if v, ok := sc.mutation.Principal(); ok {
		if err := session.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Session.principal": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Hostname(); ok {
		if err := session.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`ent: validator failed for field "Session.hostname": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Session.identifier"`)}
	}
	if v, ok := sc.mutation.Identifier(); ok {
		if err := session.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "Session.identifier": %w`, err)}
		}
	}
	if v, ok := sc.mutation.AgentIdentifier(); ok {
		if err := session.AgentIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "agentIdentifier", err: fmt.Errorf(`ent: validator failed for field "Session.agentIdentifier": %w`, err)}
		}
	}
	if v, ok := sc.mutation.HostIdentifier(); ok {
		if err := session.HostIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "hostIdentifier", err: fmt.Errorf(`ent: validator failed for field "Session.hostIdentifier": %w`, err)}
		}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: session.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Principal(); ok {
		_spec.SetField(session.FieldPrincipal, field.TypeString, value)
		_node.Principal = value
	}
	if value, ok := sc.mutation.Hostname(); ok {
		_spec.SetField(session.FieldHostname, field.TypeString, value)
		_node.Hostname = value
	}
	if value, ok := sc.mutation.Identifier(); ok {
		_spec.SetField(session.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	if value, ok := sc.mutation.AgentIdentifier(); ok {
		_spec.SetField(session.FieldAgentIdentifier, field.TypeString, value)
		_node.AgentIdentifier = value
	}
	if value, ok := sc.mutation.HostIdentifier(); ok {
		_spec.SetField(session.FieldHostIdentifier, field.TypeString, value)
		_node.HostIdentifier = value
	}
	if value, ok := sc.mutation.LastSeenAt(); ok {
		_spec.SetField(session.FieldLastSeenAt, field.TypeTime, value)
		_node.LastSeenAt = value
	}
	if nodes := sc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   session.TagsTable,
			Columns: session.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.TasksTable,
			Columns: []string{session.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
