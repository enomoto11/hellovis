// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hellovis/ent/predicate"
	"hellovis/ent/student"
	"hellovis/ent/studentcheckin"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StudentCheckinUpdate is the builder for updating StudentCheckin entities.
type StudentCheckinUpdate struct {
	config
	hooks    []Hook
	mutation *StudentCheckinMutation
}

// Where appends a list predicates to the StudentCheckinUpdate builder.
func (scu *StudentCheckinUpdate) Where(ps ...predicate.StudentCheckin) *StudentCheckinUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetUpdatedAt sets the "updated_at" field.
func (scu *StudentCheckinUpdate) SetUpdatedAt(t time.Time) *StudentCheckinUpdate {
	scu.mutation.SetUpdatedAt(t)
	return scu
}

// SetStudentID sets the "student_id" field.
func (scu *StudentCheckinUpdate) SetStudentID(u uuid.UUID) *StudentCheckinUpdate {
	scu.mutation.SetStudentID(u)
	return scu
}

// SetStudent sets the "student" edge to the Student entity.
func (scu *StudentCheckinUpdate) SetStudent(s *Student) *StudentCheckinUpdate {
	return scu.SetStudentID(s.ID)
}

// Mutation returns the StudentCheckinMutation object of the builder.
func (scu *StudentCheckinUpdate) Mutation() *StudentCheckinMutation {
	return scu.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (scu *StudentCheckinUpdate) ClearStudent() *StudentCheckinUpdate {
	scu.mutation.ClearStudent()
	return scu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *StudentCheckinUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *StudentCheckinUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *StudentCheckinUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *StudentCheckinUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *StudentCheckinUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := studentcheckin.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *StudentCheckinUpdate) check() error {
	if _, ok := scu.mutation.StudentID(); scu.mutation.StudentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "StudentCheckin.student"`)
	}
	return nil
}

func (scu *StudentCheckinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(studentcheckin.Table, studentcheckin.Columns, sqlgraph.NewFieldSpec(studentcheckin.FieldID, field.TypeUUID))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.SetField(studentcheckin.FieldUpdatedAt, field.TypeTime, value)
	}
	if scu.mutation.StudentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.StudentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studentcheckin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// StudentCheckinUpdateOne is the builder for updating a single StudentCheckin entity.
type StudentCheckinUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentCheckinMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (scuo *StudentCheckinUpdateOne) SetUpdatedAt(t time.Time) *StudentCheckinUpdateOne {
	scuo.mutation.SetUpdatedAt(t)
	return scuo
}

// SetStudentID sets the "student_id" field.
func (scuo *StudentCheckinUpdateOne) SetStudentID(u uuid.UUID) *StudentCheckinUpdateOne {
	scuo.mutation.SetStudentID(u)
	return scuo
}

// SetStudent sets the "student" edge to the Student entity.
func (scuo *StudentCheckinUpdateOne) SetStudent(s *Student) *StudentCheckinUpdateOne {
	return scuo.SetStudentID(s.ID)
}

// Mutation returns the StudentCheckinMutation object of the builder.
func (scuo *StudentCheckinUpdateOne) Mutation() *StudentCheckinMutation {
	return scuo.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (scuo *StudentCheckinUpdateOne) ClearStudent() *StudentCheckinUpdateOne {
	scuo.mutation.ClearStudent()
	return scuo
}

// Where appends a list predicates to the StudentCheckinUpdate builder.
func (scuo *StudentCheckinUpdateOne) Where(ps ...predicate.StudentCheckin) *StudentCheckinUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *StudentCheckinUpdateOne) Select(field string, fields ...string) *StudentCheckinUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated StudentCheckin entity.
func (scuo *StudentCheckinUpdateOne) Save(ctx context.Context) (*StudentCheckin, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *StudentCheckinUpdateOne) SaveX(ctx context.Context) *StudentCheckin {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *StudentCheckinUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *StudentCheckinUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *StudentCheckinUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := studentcheckin.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *StudentCheckinUpdateOne) check() error {
	if _, ok := scuo.mutation.StudentID(); scuo.mutation.StudentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "StudentCheckin.student"`)
	}
	return nil
}

func (scuo *StudentCheckinUpdateOne) sqlSave(ctx context.Context) (_node *StudentCheckin, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(studentcheckin.Table, studentcheckin.Columns, sqlgraph.NewFieldSpec(studentcheckin.FieldID, field.TypeUUID))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StudentCheckin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studentcheckin.FieldID)
		for _, f := range fields {
			if !studentcheckin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != studentcheckin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.SetField(studentcheckin.FieldUpdatedAt, field.TypeTime, value)
	}
	if scuo.mutation.StudentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.StudentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudentCheckin{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studentcheckin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
