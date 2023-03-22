// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/file"
	"github.com/kcarretto/realm/tavern/ent/tome"
)

// TomeCreate is the builder for creating a Tome entity.
type TomeCreate struct {
	config
	mutation *TomeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TomeCreate) SetCreatedAt(t time.Time) *TomeCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TomeCreate) SetNillableCreatedAt(t *time.Time) *TomeCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (tc *TomeCreate) SetLastModifiedAt(t time.Time) *TomeCreate {
	tc.mutation.SetLastModifiedAt(t)
	return tc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (tc *TomeCreate) SetNillableLastModifiedAt(t *time.Time) *TomeCreate {
	if t != nil {
		tc.SetLastModifiedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TomeCreate) SetName(s string) *TomeCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TomeCreate) SetDescription(s string) *TomeCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetParamDefs sets the "param_defs" field.
func (tc *TomeCreate) SetParamDefs(s string) *TomeCreate {
	tc.mutation.SetParamDefs(s)
	return tc
}

// SetNillableParamDefs sets the "param_defs" field if the given value is not nil.
func (tc *TomeCreate) SetNillableParamDefs(s *string) *TomeCreate {
	if s != nil {
		tc.SetParamDefs(*s)
	}
	return tc
}

// SetHash sets the "hash" field.
func (tc *TomeCreate) SetHash(s string) *TomeCreate {
	tc.mutation.SetHash(s)
	return tc
}

// SetEldritch sets the "eldritch" field.
func (tc *TomeCreate) SetEldritch(s string) *TomeCreate {
	tc.mutation.SetEldritch(s)
	return tc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (tc *TomeCreate) AddFileIDs(ids ...int) *TomeCreate {
	tc.mutation.AddFileIDs(ids...)
	return tc
}

// AddFiles adds the "files" edges to the File entity.
func (tc *TomeCreate) AddFiles(f ...*File) *TomeCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddFileIDs(ids...)
}

// Mutation returns the TomeMutation object of the builder.
func (tc *TomeCreate) Mutation() *TomeMutation {
	return tc.mutation
}

// Save creates the Tome in the database.
func (tc *TomeCreate) Save(ctx context.Context) (*Tome, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Tome, TomeMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TomeCreate) SaveX(ctx context.Context) *Tome {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TomeCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TomeCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TomeCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if tome.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized tome.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := tome.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.LastModifiedAt(); !ok {
		if tome.DefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized tome.DefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := tome.DefaultLastModifiedAt()
		tc.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TomeCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tome.created_at"`)}
	}
	if _, ok := tc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "Tome.last_modified_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Tome.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := tome.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tome.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Tome.description"`)}
	}
	if _, ok := tc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Tome.hash"`)}
	}
	if v, ok := tc.mutation.Hash(); ok {
		if err := tome.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Tome.hash": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Eldritch(); !ok {
		return &ValidationError{Name: "eldritch", err: errors.New(`ent: missing required field "Tome.eldritch"`)}
	}
	return nil
}

func (tc *TomeCreate) sqlSave(ctx context.Context) (*Tome, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TomeCreate) createSpec() (*Tome, *sqlgraph.CreateSpec) {
	var (
		_node = &Tome{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tome.Table, sqlgraph.NewFieldSpec(tome.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tome.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.LastModifiedAt(); ok {
		_spec.SetField(tome.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(tome.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(tome.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.ParamDefs(); ok {
		_spec.SetField(tome.FieldParamDefs, field.TypeString, value)
		_node.ParamDefs = value
	}
	if value, ok := tc.mutation.Hash(); ok {
		_spec.SetField(tome.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := tc.mutation.Eldritch(); ok {
		_spec.SetField(tome.FieldEldritch, field.TypeString, value)
		_node.Eldritch = value
	}
	if nodes := tc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: []string{tome.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
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

// TomeCreateBulk is the builder for creating many Tome entities in bulk.
type TomeCreateBulk struct {
	config
	builders []*TomeCreate
}

// Save creates the Tome entities in the database.
func (tcb *TomeCreateBulk) Save(ctx context.Context) ([]*Tome, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tome, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TomeMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TomeCreateBulk) SaveX(ctx context.Context) []*Tome {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TomeCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TomeCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
