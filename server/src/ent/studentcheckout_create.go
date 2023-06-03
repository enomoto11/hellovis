// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hellovis/ent/studentcheckout"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentCheckoutCreate is the builder for creating a StudentCheckout entity.
type StudentCheckoutCreate struct {
	config
	mutation *StudentCheckoutMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *StudentCheckoutCreate) SetCreatedAt(t time.Time) *StudentCheckoutCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *StudentCheckoutCreate) SetNillableCreatedAt(t *time.Time) *StudentCheckoutCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *StudentCheckoutCreate) SetUpdatedAt(t time.Time) *StudentCheckoutCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *StudentCheckoutCreate) SetNillableUpdatedAt(t *time.Time) *StudentCheckoutCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetID sets the "id" field.
func (scc *StudentCheckoutCreate) SetID(u uuid.UUID) *StudentCheckoutCreate {
	scc.mutation.SetID(u)
	return scc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scc *StudentCheckoutCreate) SetNillableID(u *uuid.UUID) *StudentCheckoutCreate {
	if u != nil {
		scc.SetID(*u)
	}
	return scc
}

// Mutation returns the StudentCheckoutMutation object of the builder.
func (scc *StudentCheckoutCreate) Mutation() *StudentCheckoutMutation {
	return scc.mutation
}

// Save creates the StudentCheckout in the database.
func (scc *StudentCheckoutCreate) Save(ctx context.Context) (*StudentCheckout, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *StudentCheckoutCreate) SaveX(ctx context.Context) *StudentCheckout {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *StudentCheckoutCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *StudentCheckoutCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *StudentCheckoutCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := studentcheckout.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := studentcheckout.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
	if _, ok := scc.mutation.ID(); !ok {
		v := studentcheckout.DefaultID()
		scc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *StudentCheckoutCreate) check() error {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "StudentCheckout.created_at"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "StudentCheckout.updated_at"`)}
	}
	return nil
}

func (scc *StudentCheckoutCreate) sqlSave(ctx context.Context) (*StudentCheckout, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *StudentCheckoutCreate) createSpec() (*StudentCheckout, *sqlgraph.CreateSpec) {
	var (
		_node = &StudentCheckout{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(studentcheckout.Table, sqlgraph.NewFieldSpec(studentcheckout.FieldID, field.TypeUUID))
	)
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.SetField(studentcheckout.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.SetField(studentcheckout.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// StudentCheckoutCreateBulk is the builder for creating many StudentCheckout entities in bulk.
type StudentCheckoutCreateBulk struct {
	config
	builders []*StudentCheckoutCreate
}

// Save creates the StudentCheckout entities in the database.
func (sccb *StudentCheckoutCreateBulk) Save(ctx context.Context) ([]*StudentCheckout, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*StudentCheckout, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentCheckoutMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *StudentCheckoutCreateBulk) SaveX(ctx context.Context) []*StudentCheckout {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *StudentCheckoutCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *StudentCheckoutCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
