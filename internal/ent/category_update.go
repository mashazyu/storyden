// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/category"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks     []Hook
	mutation  *CategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CategoryUpdate) SetUpdatedAt(t time.Time) *CategoryUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoryUpdate) SetDescription(s string) *CategoryUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDescription(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// SetColour sets the "colour" field.
func (cu *CategoryUpdate) SetColour(s string) *CategoryUpdate {
	cu.mutation.SetColour(s)
	return cu
}

// SetNillableColour sets the "colour" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableColour(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetColour(*s)
	}
	return cu
}

// SetSort sets the "sort" field.
func (cu *CategoryUpdate) SetSort(i int) *CategoryUpdate {
	cu.mutation.ResetSort()
	cu.mutation.SetSort(i)
	return cu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableSort(i *int) *CategoryUpdate {
	if i != nil {
		cu.SetSort(*i)
	}
	return cu
}

// AddSort adds i to the "sort" field.
func (cu *CategoryUpdate) AddSort(i int) *CategoryUpdate {
	cu.mutation.AddSort(i)
	return cu
}

// SetAdmin sets the "admin" field.
func (cu *CategoryUpdate) SetAdmin(b bool) *CategoryUpdate {
	cu.mutation.SetAdmin(b)
	return cu
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableAdmin(b *bool) *CategoryUpdate {
	if b != nil {
		cu.SetAdmin(*b)
	}
	return cu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cu *CategoryUpdate) AddPostIDs(ids ...xid.ID) *CategoryUpdate {
	cu.mutation.AddPostIDs(ids...)
	return cu
}

// AddPosts adds the "posts" edges to the Post entity.
func (cu *CategoryUpdate) AddPosts(p ...*Post) *CategoryUpdate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPostIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (cu *CategoryUpdate) ClearPosts() *CategoryUpdate {
	cu.mutation.ClearPosts()
	return cu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (cu *CategoryUpdate) RemovePostIDs(ids ...xid.ID) *CategoryUpdate {
	cu.mutation.RemovePostIDs(ids...)
	return cu
}

// RemovePosts removes "posts" edges to Post entities.
func (cu *CategoryUpdate) RemovePosts(p ...*Post) *CategoryUpdate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, CategoryMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CategoryUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := category.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CategoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CategoryUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.Colour(); ok {
		_spec.SetField(category.FieldColour, field.TypeString, value)
	}
	if value, ok := cu.mutation.Sort(); ok {
		_spec.SetField(category.FieldSort, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedSort(); ok {
		_spec.AddField(category.FieldSort, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Admin(); ok {
		_spec.SetField(category.FieldAdmin, field.TypeBool, value)
	}
	if cu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !cu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CategoryUpdateOne) SetUpdatedAt(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoryUpdateOne) SetDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDescription(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// SetColour sets the "colour" field.
func (cuo *CategoryUpdateOne) SetColour(s string) *CategoryUpdateOne {
	cuo.mutation.SetColour(s)
	return cuo
}

// SetNillableColour sets the "colour" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableColour(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetColour(*s)
	}
	return cuo
}

// SetSort sets the "sort" field.
func (cuo *CategoryUpdateOne) SetSort(i int) *CategoryUpdateOne {
	cuo.mutation.ResetSort()
	cuo.mutation.SetSort(i)
	return cuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableSort(i *int) *CategoryUpdateOne {
	if i != nil {
		cuo.SetSort(*i)
	}
	return cuo
}

// AddSort adds i to the "sort" field.
func (cuo *CategoryUpdateOne) AddSort(i int) *CategoryUpdateOne {
	cuo.mutation.AddSort(i)
	return cuo
}

// SetAdmin sets the "admin" field.
func (cuo *CategoryUpdateOne) SetAdmin(b bool) *CategoryUpdateOne {
	cuo.mutation.SetAdmin(b)
	return cuo
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableAdmin(b *bool) *CategoryUpdateOne {
	if b != nil {
		cuo.SetAdmin(*b)
	}
	return cuo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cuo *CategoryUpdateOne) AddPostIDs(ids ...xid.ID) *CategoryUpdateOne {
	cuo.mutation.AddPostIDs(ids...)
	return cuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (cuo *CategoryUpdateOne) AddPosts(p ...*Post) *CategoryUpdateOne {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPostIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (cuo *CategoryUpdateOne) ClearPosts() *CategoryUpdateOne {
	cuo.mutation.ClearPosts()
	return cuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (cuo *CategoryUpdateOne) RemovePostIDs(ids ...xid.ID) *CategoryUpdateOne {
	cuo.mutation.RemovePostIDs(ids...)
	return cuo
}

// RemovePosts removes "posts" edges to Post entities.
func (cuo *CategoryUpdateOne) RemovePosts(p ...*Post) *CategoryUpdateOne {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePostIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	cuo.defaults()
	return withHooks[*Category, CategoryMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CategoryUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := category.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CategoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CategoryUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Colour(); ok {
		_spec.SetField(category.FieldColour, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Sort(); ok {
		_spec.SetField(category.FieldSort, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedSort(); ok {
		_spec.AddField(category.FieldSort, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Admin(); ok {
		_spec.SetField(category.FieldAdmin, field.TypeBool, value)
	}
	if cuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !cuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.PostsTable,
			Columns: []string{category.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
