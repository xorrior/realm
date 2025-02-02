// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/internal/ent/beacon"
	"github.com/kcarretto/realm/tavern/internal/ent/tag"
	"github.com/kcarretto/realm/tavern/internal/ent/task"
)

// BeaconCreate is the builder for creating a Beacon entity.
type BeaconCreate struct {
	config
	mutation *BeaconMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BeaconCreate) SetName(s string) *BeaconCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableName(s *string) *BeaconCreate {
	if s != nil {
		bc.SetName(*s)
	}
	return bc
}

// SetPrincipal sets the "principal" field.
func (bc *BeaconCreate) SetPrincipal(s string) *BeaconCreate {
	bc.mutation.SetPrincipal(s)
	return bc
}

// SetNillablePrincipal sets the "principal" field if the given value is not nil.
func (bc *BeaconCreate) SetNillablePrincipal(s *string) *BeaconCreate {
	if s != nil {
		bc.SetPrincipal(*s)
	}
	return bc
}

// SetHostname sets the "hostname" field.
func (bc *BeaconCreate) SetHostname(s string) *BeaconCreate {
	bc.mutation.SetHostname(s)
	return bc
}

// SetNillableHostname sets the "hostname" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableHostname(s *string) *BeaconCreate {
	if s != nil {
		bc.SetHostname(*s)
	}
	return bc
}

// SetIdentifier sets the "identifier" field.
func (bc *BeaconCreate) SetIdentifier(s string) *BeaconCreate {
	bc.mutation.SetIdentifier(s)
	return bc
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableIdentifier(s *string) *BeaconCreate {
	if s != nil {
		bc.SetIdentifier(*s)
	}
	return bc
}

// SetAgentIdentifier sets the "agent_identifier" field.
func (bc *BeaconCreate) SetAgentIdentifier(s string) *BeaconCreate {
	bc.mutation.SetAgentIdentifier(s)
	return bc
}

// SetNillableAgentIdentifier sets the "agent_identifier" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableAgentIdentifier(s *string) *BeaconCreate {
	if s != nil {
		bc.SetAgentIdentifier(*s)
	}
	return bc
}

// SetHostIdentifier sets the "host_identifier" field.
func (bc *BeaconCreate) SetHostIdentifier(s string) *BeaconCreate {
	bc.mutation.SetHostIdentifier(s)
	return bc
}

// SetNillableHostIdentifier sets the "host_identifier" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableHostIdentifier(s *string) *BeaconCreate {
	if s != nil {
		bc.SetHostIdentifier(*s)
	}
	return bc
}

// SetHostPrimaryIP sets the "host_primary_ip" field.
func (bc *BeaconCreate) SetHostPrimaryIP(s string) *BeaconCreate {
	bc.mutation.SetHostPrimaryIP(s)
	return bc
}

// SetNillableHostPrimaryIP sets the "host_primary_ip" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableHostPrimaryIP(s *string) *BeaconCreate {
	if s != nil {
		bc.SetHostPrimaryIP(*s)
	}
	return bc
}

// SetHostPlatform sets the "host_platform" field.
func (bc *BeaconCreate) SetHostPlatform(bp beacon.HostPlatform) *BeaconCreate {
	bc.mutation.SetHostPlatform(bp)
	return bc
}

// SetNillableHostPlatform sets the "host_platform" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableHostPlatform(bp *beacon.HostPlatform) *BeaconCreate {
	if bp != nil {
		bc.SetHostPlatform(*bp)
	}
	return bc
}

// SetLastSeenAt sets the "last_seen_at" field.
func (bc *BeaconCreate) SetLastSeenAt(t time.Time) *BeaconCreate {
	bc.mutation.SetLastSeenAt(t)
	return bc
}

// SetNillableLastSeenAt sets the "last_seen_at" field if the given value is not nil.
func (bc *BeaconCreate) SetNillableLastSeenAt(t *time.Time) *BeaconCreate {
	if t != nil {
		bc.SetLastSeenAt(*t)
	}
	return bc
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (bc *BeaconCreate) AddTagIDs(ids ...int) *BeaconCreate {
	bc.mutation.AddTagIDs(ids...)
	return bc
}

// AddTags adds the "tags" edges to the Tag entity.
func (bc *BeaconCreate) AddTags(t ...*Tag) *BeaconCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTagIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (bc *BeaconCreate) AddTaskIDs(ids ...int) *BeaconCreate {
	bc.mutation.AddTaskIDs(ids...)
	return bc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (bc *BeaconCreate) AddTasks(t ...*Task) *BeaconCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTaskIDs(ids...)
}

// Mutation returns the BeaconMutation object of the builder.
func (bc *BeaconCreate) Mutation() *BeaconMutation {
	return bc.mutation
}

// Save creates the Beacon in the database.
func (bc *BeaconCreate) Save(ctx context.Context) (*Beacon, error) {
	bc.defaults()
	return withHooks[*Beacon, BeaconMutation](ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BeaconCreate) SaveX(ctx context.Context) *Beacon {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BeaconCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BeaconCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BeaconCreate) defaults() {
	if _, ok := bc.mutation.Name(); !ok {
		v := beacon.DefaultName()
		bc.mutation.SetName(v)
	}
	if _, ok := bc.mutation.Identifier(); !ok {
		v := beacon.DefaultIdentifier()
		bc.mutation.SetIdentifier(v)
	}
	if _, ok := bc.mutation.HostPlatform(); !ok {
		v := beacon.DefaultHostPlatform
		bc.mutation.SetHostPlatform(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BeaconCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Beacon.name"`)}
	}
	if v, ok := bc.mutation.Name(); ok {
		if err := beacon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Beacon.name": %w`, err)}
		}
	}
	if v, ok := bc.mutation.Principal(); ok {
		if err := beacon.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Beacon.principal": %w`, err)}
		}
	}
	if v, ok := bc.mutation.Hostname(); ok {
		if err := beacon.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`ent: validator failed for field "Beacon.hostname": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Beacon.identifier"`)}
	}
	if v, ok := bc.mutation.Identifier(); ok {
		if err := beacon.IdentifierValidator(v); err != nil {
			return &ValidationError{Name: "identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.identifier": %w`, err)}
		}
	}
	if v, ok := bc.mutation.AgentIdentifier(); ok {
		if err := beacon.AgentIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "agent_identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.agent_identifier": %w`, err)}
		}
	}
	if v, ok := bc.mutation.HostIdentifier(); ok {
		if err := beacon.HostIdentifierValidator(v); err != nil {
			return &ValidationError{Name: "host_identifier", err: fmt.Errorf(`ent: validator failed for field "Beacon.host_identifier": %w`, err)}
		}
	}
	if _, ok := bc.mutation.HostPlatform(); !ok {
		return &ValidationError{Name: "host_platform", err: errors.New(`ent: missing required field "Beacon.host_platform"`)}
	}
	if v, ok := bc.mutation.HostPlatform(); ok {
		if err := beacon.HostPlatformValidator(v); err != nil {
			return &ValidationError{Name: "host_platform", err: fmt.Errorf(`ent: validator failed for field "Beacon.host_platform": %w`, err)}
		}
	}
	return nil
}

func (bc *BeaconCreate) sqlSave(ctx context.Context) (*Beacon, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BeaconCreate) createSpec() (*Beacon, *sqlgraph.CreateSpec) {
	var (
		_node = &Beacon{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(beacon.Table, sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Name(); ok {
		_spec.SetField(beacon.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bc.mutation.Principal(); ok {
		_spec.SetField(beacon.FieldPrincipal, field.TypeString, value)
		_node.Principal = value
	}
	if value, ok := bc.mutation.Hostname(); ok {
		_spec.SetField(beacon.FieldHostname, field.TypeString, value)
		_node.Hostname = value
	}
	if value, ok := bc.mutation.Identifier(); ok {
		_spec.SetField(beacon.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	if value, ok := bc.mutation.AgentIdentifier(); ok {
		_spec.SetField(beacon.FieldAgentIdentifier, field.TypeString, value)
		_node.AgentIdentifier = value
	}
	if value, ok := bc.mutation.HostIdentifier(); ok {
		_spec.SetField(beacon.FieldHostIdentifier, field.TypeString, value)
		_node.HostIdentifier = value
	}
	if value, ok := bc.mutation.HostPrimaryIP(); ok {
		_spec.SetField(beacon.FieldHostPrimaryIP, field.TypeString, value)
		_node.HostPrimaryIP = value
	}
	if value, ok := bc.mutation.HostPlatform(); ok {
		_spec.SetField(beacon.FieldHostPlatform, field.TypeEnum, value)
		_node.HostPlatform = value
	}
	if value, ok := bc.mutation.LastSeenAt(); ok {
		_spec.SetField(beacon.FieldLastSeenAt, field.TypeTime, value)
		_node.LastSeenAt = value
	}
	if nodes := bc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   beacon.TagsTable,
			Columns: beacon.TagsPrimaryKey,
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
	if nodes := bc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   beacon.TasksTable,
			Columns: []string{beacon.TasksColumn},
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

// BeaconCreateBulk is the builder for creating many Beacon entities in bulk.
type BeaconCreateBulk struct {
	config
	builders []*BeaconCreate
}

// Save creates the Beacon entities in the database.
func (bcb *BeaconCreateBulk) Save(ctx context.Context) ([]*Beacon, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Beacon, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BeaconMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BeaconCreateBulk) SaveX(ctx context.Context) []*Beacon {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BeaconCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BeaconCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
