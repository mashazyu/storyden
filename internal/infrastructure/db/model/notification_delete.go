// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/notification"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/predicate"
)

// NotificationDelete is the builder for deleting a Notification entity.
type NotificationDelete struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationDelete builder.
func (nd *NotificationDelete) Where(ps ...predicate.Notification) *NotificationDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NotificationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nd.hooks) == 0 {
		affected, err = nd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NotificationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nd.mutation = mutation
			affected, err = nd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nd.hooks) - 1; i >= 0; i-- {
			if nd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = nd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NotificationDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NotificationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: notification.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeBytes,
				Column: notification.FieldID,
			},
		},
	}
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// NotificationDeleteOne is the builder for deleting a single Notification entity.
type NotificationDeleteOne struct {
	nd *NotificationDelete
}

// Exec executes the deletion query.
func (ndo *NotificationDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notification.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NotificationDeleteOne) ExecX(ctx context.Context) {
	ndo.nd.ExecX(ctx)
}
