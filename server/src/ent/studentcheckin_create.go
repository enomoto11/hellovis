// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hellovis/ent/student"
	"hellovis/ent/studentcheckin"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentCheckinCreate is the builder for creating a StudentCheckin entity.
type StudentCheckinCreate struct {
	config
	mutation *StudentCheckinMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *StudentCheckinCreate) SetCreatedAt(t time.Time) *StudentCheckinCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *StudentCheckinCreate) SetNillableCreatedAt(t *time.Time) *StudentCheckinCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *StudentCheckinCreate) SetUpdatedAt(t time.Time) *StudentCheckinCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *StudentCheckinCreate) SetNillableUpdatedAt(t *time.Time) *StudentCheckinCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetStudentID sets the "student_id" field.
func (scc *StudentCheckinCreate) SetStudentID(u uuid.UUID) *StudentCheckinCreate {
	scc.mutation.SetStudentID(u)
	return scc
}

// SetCheckinAt sets the "checkin_at" field.
func (scc *StudentCheckinCreate) SetCheckinAt(t time.Time) *StudentCheckinCreate {
	scc.mutation.SetCheckinAt(t)
	return scc
}

// SetNillableCheckinAt sets the "checkin_at" field if the given value is not nil.
func (scc *StudentCheckinCreate) SetNillableCheckinAt(t *time.Time) *StudentCheckinCreate {
	if t != nil {
		scc.SetCheckinAt(*t)
	}
	return scc
}

// SetID sets the "id" field.
func (scc *StudentCheckinCreate) SetID(u uuid.UUID) *StudentCheckinCreate {
	scc.mutation.SetID(u)
	return scc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scc *StudentCheckinCreate) SetNillableID(u *uuid.UUID) *StudentCheckinCreate {
	if u != nil {
		scc.SetID(*u)
	}
	return scc
}

// SetStudent sets the "student" edge to the Student entity.
func (scc *StudentCheckinCreate) SetStudent(s *Student) *StudentCheckinCreate {
	return scc.SetStudentID(s.ID)
}

// Mutation returns the StudentCheckinMutation object of the builder.
func (scc *StudentCheckinCreate) Mutation() *StudentCheckinMutation {
	return scc.mutation
}

// Save creates the StudentCheckin in the database.
func (scc *StudentCheckinCreate) Save(ctx context.Context) (*StudentCheckin, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *StudentCheckinCreate) SaveX(ctx context.Context) *StudentCheckin {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *StudentCheckinCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *StudentCheckinCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *StudentCheckinCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := studentcheckin.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := studentcheckin.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
	if _, ok := scc.mutation.CheckinAt(); !ok {
		v := studentcheckin.DefaultCheckinAt()
		scc.mutation.SetCheckinAt(v)
	}
	if _, ok := scc.mutation.ID(); !ok {
		v := studentcheckin.DefaultID()
		scc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *StudentCheckinCreate) check() error {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "StudentCheckin.created_at"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "StudentCheckin.updated_at"`)}
	}
	if _, ok := scc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student_id", err: errors.New(`ent: missing required field "StudentCheckin.student_id"`)}
	}
	if _, ok := scc.mutation.CheckinAt(); !ok {
		return &ValidationError{Name: "checkin_at", err: errors.New(`ent: missing required field "StudentCheckin.checkin_at"`)}
	}
	if _, ok := scc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student", err: errors.New(`ent: missing required edge "StudentCheckin.student"`)}
	}
	return nil
}

func (scc *StudentCheckinCreate) sqlSave(ctx context.Context) (*StudentCheckin, error) {
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

func (scc *StudentCheckinCreate) createSpec() (*StudentCheckin, *sqlgraph.CreateSpec) {
	var (
		_node = &StudentCheckin{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(studentcheckin.Table, sqlgraph.NewFieldSpec(studentcheckin.FieldID, field.TypeUUID))
	)
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.SetField(studentcheckin.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.SetField(studentcheckin.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.CheckinAt(); ok {
		_spec.SetField(studentcheckin.FieldCheckinAt, field.TypeTime, value)
		_node.CheckinAt = value
	}
	if nodes := scc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studentcheckin.StudentTable,
			Columns: []string{studentcheckin.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.StudentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudentCheckinCreateBulk is the builder for creating many StudentCheckin entities in bulk.
type StudentCheckinCreateBulk struct {
	config
	builders []*StudentCheckinCreate
}

// Save creates the StudentCheckin entities in the database.
func (sccb *StudentCheckinCreateBulk) Save(ctx context.Context) ([]*StudentCheckin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*StudentCheckin, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentCheckinMutation)
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
func (sccb *StudentCheckinCreateBulk) SaveX(ctx context.Context) []*StudentCheckin {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *StudentCheckinCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *StudentCheckinCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
