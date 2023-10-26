// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/item"
	"github.com/Southclaws/storyden/internal/ent/predicate"
)

// ItemDelete is the builder for deleting a Item entity.
type ItemDelete struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemDelete builder.
func (id *ItemDelete) Where(ps ...predicate.Item) *ItemDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *ItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *ItemDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *ItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(item.Table, sqlgraph.NewFieldSpec(item.FieldID, field.TypeString))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// ItemDeleteOne is the builder for deleting a single Item entity.
type ItemDeleteOne struct {
	id *ItemDelete
}

// Where appends a list predicates to the ItemDelete builder.
func (ido *ItemDeleteOne) Where(ps ...predicate.Item) *ItemDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *ItemDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{item.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *ItemDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}