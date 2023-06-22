// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/authentication"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// AuthenticationUpdate is the builder for updating Authentication entities.
type AuthenticationUpdate struct {
	config
	hooks     []Hook
	mutation  *AuthenticationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AuthenticationUpdate builder.
func (au *AuthenticationUpdate) Where(ps ...predicate.Authentication) *AuthenticationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetService sets the "service" field.
func (au *AuthenticationUpdate) SetService(s string) *AuthenticationUpdate {
	au.mutation.SetService(s)
	return au
}

// SetIdentifier sets the "identifier" field.
func (au *AuthenticationUpdate) SetIdentifier(s string) *AuthenticationUpdate {
	au.mutation.SetIdentifier(s)
	return au
}

// SetToken sets the "token" field.
func (au *AuthenticationUpdate) SetToken(s string) *AuthenticationUpdate {
	au.mutation.SetToken(s)
	return au
}

// SetMetadata sets the "metadata" field.
func (au *AuthenticationUpdate) SetMetadata(m map[string]interface{}) *AuthenticationUpdate {
	au.mutation.SetMetadata(m)
	return au
}

// ClearMetadata clears the value of the "metadata" field.
func (au *AuthenticationUpdate) ClearMetadata() *AuthenticationUpdate {
	au.mutation.ClearMetadata()
	return au
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (au *AuthenticationUpdate) SetAccountID(id xid.ID) *AuthenticationUpdate {
	au.mutation.SetAccountID(id)
	return au
}

// SetAccount sets the "account" edge to the Account entity.
func (au *AuthenticationUpdate) SetAccount(a *Account) *AuthenticationUpdate {
	return au.SetAccountID(a.ID)
}

// Mutation returns the AuthenticationMutation object of the builder.
func (au *AuthenticationUpdate) Mutation() *AuthenticationMutation {
	return au.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (au *AuthenticationUpdate) ClearAccount() *AuthenticationUpdate {
	au.mutation.ClearAccount()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuthenticationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AuthenticationMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthenticationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthenticationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthenticationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AuthenticationUpdate) check() error {
	if v, ok := au.mutation.Service(); ok {
		if err := authentication.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`ent: validator failed for field "Authentication.service": %w`, err)}
		}
	}
	if v, ok := au.mutation.Token(); ok {
		if err := authentication.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Authentication.token": %w`, err)}
		}
	}
	if _, ok := au.mutation.AccountID(); au.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Authentication.account"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AuthenticationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AuthenticationUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AuthenticationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(authentication.Table, authentication.Columns, sqlgraph.NewFieldSpec(authentication.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Service(); ok {
		_spec.SetField(authentication.FieldService, field.TypeString, value)
	}
	if value, ok := au.mutation.Identifier(); ok {
		_spec.SetField(authentication.FieldIdentifier, field.TypeString, value)
	}
	if value, ok := au.mutation.Token(); ok {
		_spec.SetField(authentication.FieldToken, field.TypeString, value)
	}
	if value, ok := au.mutation.Metadata(); ok {
		_spec.SetField(authentication.FieldMetadata, field.TypeJSON, value)
	}
	if au.mutation.MetadataCleared() {
		_spec.ClearField(authentication.FieldMetadata, field.TypeJSON)
	}
	if au.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authentication.AccountTable,
			Columns: []string{authentication.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authentication.AccountTable,
			Columns: []string{authentication.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authentication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AuthenticationUpdateOne is the builder for updating a single Authentication entity.
type AuthenticationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AuthenticationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetService sets the "service" field.
func (auo *AuthenticationUpdateOne) SetService(s string) *AuthenticationUpdateOne {
	auo.mutation.SetService(s)
	return auo
}

// SetIdentifier sets the "identifier" field.
func (auo *AuthenticationUpdateOne) SetIdentifier(s string) *AuthenticationUpdateOne {
	auo.mutation.SetIdentifier(s)
	return auo
}

// SetToken sets the "token" field.
func (auo *AuthenticationUpdateOne) SetToken(s string) *AuthenticationUpdateOne {
	auo.mutation.SetToken(s)
	return auo
}

// SetMetadata sets the "metadata" field.
func (auo *AuthenticationUpdateOne) SetMetadata(m map[string]interface{}) *AuthenticationUpdateOne {
	auo.mutation.SetMetadata(m)
	return auo
}

// ClearMetadata clears the value of the "metadata" field.
func (auo *AuthenticationUpdateOne) ClearMetadata() *AuthenticationUpdateOne {
	auo.mutation.ClearMetadata()
	return auo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (auo *AuthenticationUpdateOne) SetAccountID(id xid.ID) *AuthenticationUpdateOne {
	auo.mutation.SetAccountID(id)
	return auo
}

// SetAccount sets the "account" edge to the Account entity.
func (auo *AuthenticationUpdateOne) SetAccount(a *Account) *AuthenticationUpdateOne {
	return auo.SetAccountID(a.ID)
}

// Mutation returns the AuthenticationMutation object of the builder.
func (auo *AuthenticationUpdateOne) Mutation() *AuthenticationMutation {
	return auo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (auo *AuthenticationUpdateOne) ClearAccount() *AuthenticationUpdateOne {
	auo.mutation.ClearAccount()
	return auo
}

// Where appends a list predicates to the AuthenticationUpdate builder.
func (auo *AuthenticationUpdateOne) Where(ps ...predicate.Authentication) *AuthenticationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuthenticationUpdateOne) Select(field string, fields ...string) *AuthenticationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Authentication entity.
func (auo *AuthenticationUpdateOne) Save(ctx context.Context) (*Authentication, error) {
	return withHooks[*Authentication, AuthenticationMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthenticationUpdateOne) SaveX(ctx context.Context) *Authentication {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuthenticationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthenticationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AuthenticationUpdateOne) check() error {
	if v, ok := auo.mutation.Service(); ok {
		if err := authentication.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`ent: validator failed for field "Authentication.service": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Token(); ok {
		if err := authentication.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "Authentication.token": %w`, err)}
		}
	}
	if _, ok := auo.mutation.AccountID(); auo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Authentication.account"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AuthenticationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AuthenticationUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AuthenticationUpdateOne) sqlSave(ctx context.Context) (_node *Authentication, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(authentication.Table, authentication.Columns, sqlgraph.NewFieldSpec(authentication.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Authentication.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authentication.FieldID)
		for _, f := range fields {
			if !authentication.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authentication.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Service(); ok {
		_spec.SetField(authentication.FieldService, field.TypeString, value)
	}
	if value, ok := auo.mutation.Identifier(); ok {
		_spec.SetField(authentication.FieldIdentifier, field.TypeString, value)
	}
	if value, ok := auo.mutation.Token(); ok {
		_spec.SetField(authentication.FieldToken, field.TypeString, value)
	}
	if value, ok := auo.mutation.Metadata(); ok {
		_spec.SetField(authentication.FieldMetadata, field.TypeJSON, value)
	}
	if auo.mutation.MetadataCleared() {
		_spec.ClearField(authentication.FieldMetadata, field.TypeJSON)
	}
	if auo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authentication.AccountTable,
			Columns: []string{authentication.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authentication.AccountTable,
			Columns: []string{authentication.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Authentication{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authentication.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
