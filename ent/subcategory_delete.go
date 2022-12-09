// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"shopular/ent/predicate"
	"shopular/ent/subcategory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubCategoryDelete is the builder for deleting a SubCategory entity.
type SubCategoryDelete struct {
	config
	hooks    []Hook
	mutation *SubCategoryMutation
}

// Where appends a list predicates to the SubCategoryDelete builder.
func (scd *SubCategoryDelete) Where(ps ...predicate.SubCategory) *SubCategoryDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SubCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scd.hooks) == 0 {
		affected, err = scd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scd.mutation = mutation
			affected, err = scd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scd.hooks) - 1; i >= 0; i-- {
			if scd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SubCategoryDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SubCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: subcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subcategory.FieldID,
			},
		},
	}
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SubCategoryDeleteOne is the builder for deleting a single SubCategory entity.
type SubCategoryDeleteOne struct {
	scd *SubCategoryDelete
}

// Exec executes the deletion query.
func (scdo *SubCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SubCategoryDeleteOne) ExecX(ctx context.Context) {
	scdo.scd.ExecX(ctx)
}
