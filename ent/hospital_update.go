// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"task_manager/ent/hospital"
	"task_manager/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HospitalUpdate is the builder for updating Hospital entities.
type HospitalUpdate struct {
	config
	hooks    []Hook
	mutation *HospitalMutation
}

// Where appends a list predicates to the HospitalUpdate builder.
func (hu *HospitalUpdate) Where(ps ...predicate.Hospital) *HospitalUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetName sets the "name" field.
func (hu *HospitalUpdate) SetName(s string) *HospitalUpdate {
	hu.mutation.SetName(s)
	return hu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (hu *HospitalUpdate) SetNillableName(s *string) *HospitalUpdate {
	if s != nil {
		hu.SetName(*s)
	}
	return hu
}

// SetAddress sets the "address" field.
func (hu *HospitalUpdate) SetAddress(s string) *HospitalUpdate {
	hu.mutation.SetAddress(s)
	return hu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (hu *HospitalUpdate) SetNillableAddress(s *string) *HospitalUpdate {
	if s != nil {
		hu.SetAddress(*s)
	}
	return hu
}

// SetCreatedAt sets the "created_at" field.
func (hu *HospitalUpdate) SetCreatedAt(t time.Time) *HospitalUpdate {
	hu.mutation.SetCreatedAt(t)
	return hu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hu *HospitalUpdate) SetNillableCreatedAt(t *time.Time) *HospitalUpdate {
	if t != nil {
		hu.SetCreatedAt(*t)
	}
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HospitalUpdate) SetUpdatedAt(t time.Time) *HospitalUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hu *HospitalUpdate) SetNillableUpdatedAt(t *time.Time) *HospitalUpdate {
	if t != nil {
		hu.SetUpdatedAt(*t)
	}
	return hu
}

// Mutation returns the HospitalMutation object of the builder.
func (hu *HospitalUpdate) Mutation() *HospitalMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HospitalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		if err = hu.check(); err != nil {
			return 0, err
		}
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HospitalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hu.check(); err != nil {
				return 0, err
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			if hu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HospitalUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HospitalUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HospitalUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HospitalUpdate) check() error {
	if v, ok := hu.mutation.Name(); ok {
		if err := hospital.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Hospital.name": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Address(); ok {
		if err := hospital.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Hospital.address": %w`, err)}
		}
	}
	return nil
}

func (hu *HospitalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hospital.Table,
			Columns: hospital.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hospital.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hospital.FieldName,
		})
	}
	if value, ok := hu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hospital.FieldAddress,
		})
	}
	if value, ok := hu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hospital.FieldCreatedAt,
		})
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hospital.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hospital.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// HospitalUpdateOne is the builder for updating a single Hospital entity.
type HospitalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HospitalMutation
}

// SetName sets the "name" field.
func (huo *HospitalUpdateOne) SetName(s string) *HospitalUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (huo *HospitalUpdateOne) SetNillableName(s *string) *HospitalUpdateOne {
	if s != nil {
		huo.SetName(*s)
	}
	return huo
}

// SetAddress sets the "address" field.
func (huo *HospitalUpdateOne) SetAddress(s string) *HospitalUpdateOne {
	huo.mutation.SetAddress(s)
	return huo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (huo *HospitalUpdateOne) SetNillableAddress(s *string) *HospitalUpdateOne {
	if s != nil {
		huo.SetAddress(*s)
	}
	return huo
}

// SetCreatedAt sets the "created_at" field.
func (huo *HospitalUpdateOne) SetCreatedAt(t time.Time) *HospitalUpdateOne {
	huo.mutation.SetCreatedAt(t)
	return huo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (huo *HospitalUpdateOne) SetNillableCreatedAt(t *time.Time) *HospitalUpdateOne {
	if t != nil {
		huo.SetCreatedAt(*t)
	}
	return huo
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HospitalUpdateOne) SetUpdatedAt(t time.Time) *HospitalUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (huo *HospitalUpdateOne) SetNillableUpdatedAt(t *time.Time) *HospitalUpdateOne {
	if t != nil {
		huo.SetUpdatedAt(*t)
	}
	return huo
}

// Mutation returns the HospitalMutation object of the builder.
func (huo *HospitalUpdateOne) Mutation() *HospitalMutation {
	return huo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HospitalUpdateOne) Select(field string, fields ...string) *HospitalUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Hospital entity.
func (huo *HospitalUpdateOne) Save(ctx context.Context) (*Hospital, error) {
	var (
		err  error
		node *Hospital
	)
	if len(huo.hooks) == 0 {
		if err = huo.check(); err != nil {
			return nil, err
		}
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HospitalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = huo.check(); err != nil {
				return nil, err
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			if huo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = huo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, huo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Hospital)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HospitalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HospitalUpdateOne) SaveX(ctx context.Context) *Hospital {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HospitalUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HospitalUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HospitalUpdateOne) check() error {
	if v, ok := huo.mutation.Name(); ok {
		if err := hospital.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Hospital.name": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Address(); ok {
		if err := hospital.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Hospital.address": %w`, err)}
		}
	}
	return nil
}

func (huo *HospitalUpdateOne) sqlSave(ctx context.Context) (_node *Hospital, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hospital.Table,
			Columns: hospital.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hospital.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Hospital.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hospital.FieldID)
		for _, f := range fields {
			if !hospital.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hospital.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hospital.FieldName,
		})
	}
	if value, ok := huo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hospital.FieldAddress,
		})
	}
	if value, ok := huo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hospital.FieldCreatedAt,
		})
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hospital.FieldUpdatedAt,
		})
	}
	_node = &Hospital{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hospital.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
