// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/cluster"
	"github.com/Southclaws/storyden/internal/ent/item"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/rs/xid"
)

// AssetCreate is the builder for creating a Asset entity.
type AssetCreate struct {
	config
	mutation *AssetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *AssetCreate) SetCreatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableCreatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AssetCreate) SetUpdatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableUpdatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetURL sets the "url" field.
func (ac *AssetCreate) SetURL(s string) *AssetCreate {
	ac.mutation.SetURL(s)
	return ac
}

// SetMimetype sets the "mimetype" field.
func (ac *AssetCreate) SetMimetype(s string) *AssetCreate {
	ac.mutation.SetMimetype(s)
	return ac
}

// SetWidth sets the "width" field.
func (ac *AssetCreate) SetWidth(i int) *AssetCreate {
	ac.mutation.SetWidth(i)
	return ac
}

// SetHeight sets the "height" field.
func (ac *AssetCreate) SetHeight(i int) *AssetCreate {
	ac.mutation.SetHeight(i)
	return ac
}

// SetAccountID sets the "account_id" field.
func (ac *AssetCreate) SetAccountID(x xid.ID) *AssetCreate {
	ac.mutation.SetAccountID(x)
	return ac
}

// SetID sets the "id" field.
func (ac *AssetCreate) SetID(s string) *AssetCreate {
	ac.mutation.SetID(s)
	return ac
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (ac *AssetCreate) AddPostIDs(ids ...xid.ID) *AssetCreate {
	ac.mutation.AddPostIDs(ids...)
	return ac
}

// AddPosts adds the "posts" edges to the Post entity.
func (ac *AssetCreate) AddPosts(p ...*Post) *AssetCreate {
	ids := make([]xid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPostIDs(ids...)
}

// AddClusterIDs adds the "clusters" edge to the Cluster entity by IDs.
func (ac *AssetCreate) AddClusterIDs(ids ...xid.ID) *AssetCreate {
	ac.mutation.AddClusterIDs(ids...)
	return ac
}

// AddClusters adds the "clusters" edges to the Cluster entity.
func (ac *AssetCreate) AddClusters(c ...*Cluster) *AssetCreate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddClusterIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (ac *AssetCreate) AddItemIDs(ids ...xid.ID) *AssetCreate {
	ac.mutation.AddItemIDs(ids...)
	return ac
}

// AddItems adds the "items" edges to the Item entity.
func (ac *AssetCreate) AddItems(i ...*Item) *AssetCreate {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ac.AddItemIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (ac *AssetCreate) SetOwnerID(id xid.ID) *AssetCreate {
	ac.mutation.SetOwnerID(id)
	return ac
}

// SetOwner sets the "owner" edge to the Account entity.
func (ac *AssetCreate) SetOwner(a *Account) *AssetCreate {
	return ac.SetOwnerID(a.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (ac *AssetCreate) Mutation() *AssetMutation {
	return ac.mutation
}

// Save creates the Asset in the database.
func (ac *AssetCreate) Save(ctx context.Context) (*Asset, error) {
	ac.defaults()
	return withHooks[*Asset, AssetMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AssetCreate) SaveX(ctx context.Context) *Asset {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AssetCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AssetCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AssetCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := asset.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := asset.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AssetCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Asset.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Asset.updated_at"`)}
	}
	if _, ok := ac.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Asset.url"`)}
	}
	if _, ok := ac.mutation.Mimetype(); !ok {
		return &ValidationError{Name: "mimetype", err: errors.New(`ent: missing required field "Asset.mimetype"`)}
	}
	if _, ok := ac.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Asset.width"`)}
	}
	if _, ok := ac.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Asset.height"`)}
	}
	if _, ok := ac.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Asset.account_id"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := asset.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Asset.id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Asset.owner"`)}
	}
	return nil
}

func (ac *AssetCreate) sqlSave(ctx context.Context) (*Asset, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Asset.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AssetCreate) createSpec() (*Asset, *sqlgraph.CreateSpec) {
	var (
		_node = &Asset{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(asset.Table, sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(asset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(asset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.URL(); ok {
		_spec.SetField(asset.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ac.mutation.Mimetype(); ok {
		_spec.SetField(asset.FieldMimetype, field.TypeString, value)
		_node.Mimetype = value
	}
	if value, ok := ac.mutation.Width(); ok {
		_spec.SetField(asset.FieldWidth, field.TypeInt, value)
		_node.Width = value
	}
	if value, ok := ac.mutation.Height(); ok {
		_spec.SetField(asset.FieldHeight, field.TypeInt, value)
		_node.Height = value
	}
	if nodes := ac.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.PostsTable,
			Columns: asset.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ClustersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.ClustersTable,
			Columns: asset.ClustersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asset.ItemsTable,
			Columns: asset.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.OwnerTable,
			Columns: []string{asset.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Asset.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AssetUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ac *AssetCreate) OnConflict(opts ...sql.ConflictOption) *AssetUpsertOne {
	ac.conflict = opts
	return &AssetUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AssetCreate) OnConflictColumns(columns ...string) *AssetUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AssetUpsertOne{
		create: ac,
	}
}

type (
	// AssetUpsertOne is the builder for "upsert"-ing
	//  one Asset node.
	AssetUpsertOne struct {
		create *AssetCreate
	}

	// AssetUpsert is the "OnConflict" setter.
	AssetUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AssetUpsert) SetUpdatedAt(v time.Time) *AssetUpsert {
	u.Set(asset.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssetUpsert) UpdateUpdatedAt() *AssetUpsert {
	u.SetExcluded(asset.FieldUpdatedAt)
	return u
}

// SetURL sets the "url" field.
func (u *AssetUpsert) SetURL(v string) *AssetUpsert {
	u.Set(asset.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *AssetUpsert) UpdateURL() *AssetUpsert {
	u.SetExcluded(asset.FieldURL)
	return u
}

// SetMimetype sets the "mimetype" field.
func (u *AssetUpsert) SetMimetype(v string) *AssetUpsert {
	u.Set(asset.FieldMimetype, v)
	return u
}

// UpdateMimetype sets the "mimetype" field to the value that was provided on create.
func (u *AssetUpsert) UpdateMimetype() *AssetUpsert {
	u.SetExcluded(asset.FieldMimetype)
	return u
}

// SetWidth sets the "width" field.
func (u *AssetUpsert) SetWidth(v int) *AssetUpsert {
	u.Set(asset.FieldWidth, v)
	return u
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *AssetUpsert) UpdateWidth() *AssetUpsert {
	u.SetExcluded(asset.FieldWidth)
	return u
}

// AddWidth adds v to the "width" field.
func (u *AssetUpsert) AddWidth(v int) *AssetUpsert {
	u.Add(asset.FieldWidth, v)
	return u
}

// SetHeight sets the "height" field.
func (u *AssetUpsert) SetHeight(v int) *AssetUpsert {
	u.Set(asset.FieldHeight, v)
	return u
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *AssetUpsert) UpdateHeight() *AssetUpsert {
	u.SetExcluded(asset.FieldHeight)
	return u
}

// AddHeight adds v to the "height" field.
func (u *AssetUpsert) AddHeight(v int) *AssetUpsert {
	u.Add(asset.FieldHeight, v)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *AssetUpsert) SetAccountID(v xid.ID) *AssetUpsert {
	u.Set(asset.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *AssetUpsert) UpdateAccountID() *AssetUpsert {
	u.SetExcluded(asset.FieldAccountID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(asset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AssetUpsertOne) UpdateNewValues() *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(asset.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(asset.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Asset.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AssetUpsertOne) Ignore() *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AssetUpsertOne) DoNothing() *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AssetCreate.OnConflict
// documentation for more info.
func (u *AssetUpsertOne) Update(set func(*AssetUpsert)) *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AssetUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssetUpsertOne) SetUpdatedAt(v time.Time) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateUpdatedAt() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetURL sets the "url" field.
func (u *AssetUpsertOne) SetURL(v string) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateURL() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateURL()
	})
}

// SetMimetype sets the "mimetype" field.
func (u *AssetUpsertOne) SetMimetype(v string) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetMimetype(v)
	})
}

// UpdateMimetype sets the "mimetype" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateMimetype() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateMimetype()
	})
}

// SetWidth sets the "width" field.
func (u *AssetUpsertOne) SetWidth(v int) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetWidth(v)
	})
}

// AddWidth adds v to the "width" field.
func (u *AssetUpsertOne) AddWidth(v int) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.AddWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateWidth() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateWidth()
	})
}

// SetHeight sets the "height" field.
func (u *AssetUpsertOne) SetHeight(v int) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetHeight(v)
	})
}

// AddHeight adds v to the "height" field.
func (u *AssetUpsertOne) AddHeight(v int) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.AddHeight(v)
	})
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateHeight() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateHeight()
	})
}

// SetAccountID sets the "account_id" field.
func (u *AssetUpsertOne) SetAccountID(v xid.ID) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateAccountID() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateAccountID()
	})
}

// Exec executes the query.
func (u *AssetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AssetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AssetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AssetUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AssetUpsertOne.ID is not supported by MySQL driver. Use AssetUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AssetUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AssetCreateBulk is the builder for creating many Asset entities in bulk.
type AssetCreateBulk struct {
	config
	builders []*AssetCreate
	conflict []sql.ConflictOption
}

// Save creates the Asset entities in the database.
func (acb *AssetCreateBulk) Save(ctx context.Context) ([]*Asset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Asset, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AssetCreateBulk) SaveX(ctx context.Context) []*Asset {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AssetCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AssetCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Asset.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AssetUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (acb *AssetCreateBulk) OnConflict(opts ...sql.ConflictOption) *AssetUpsertBulk {
	acb.conflict = opts
	return &AssetUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AssetCreateBulk) OnConflictColumns(columns ...string) *AssetUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AssetUpsertBulk{
		create: acb,
	}
}

// AssetUpsertBulk is the builder for "upsert"-ing
// a bulk of Asset nodes.
type AssetUpsertBulk struct {
	create *AssetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(asset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AssetUpsertBulk) UpdateNewValues() *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(asset.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(asset.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AssetUpsertBulk) Ignore() *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AssetUpsertBulk) DoNothing() *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AssetCreateBulk.OnConflict
// documentation for more info.
func (u *AssetUpsertBulk) Update(set func(*AssetUpsert)) *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AssetUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssetUpsertBulk) SetUpdatedAt(v time.Time) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateUpdatedAt() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetURL sets the "url" field.
func (u *AssetUpsertBulk) SetURL(v string) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateURL() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateURL()
	})
}

// SetMimetype sets the "mimetype" field.
func (u *AssetUpsertBulk) SetMimetype(v string) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetMimetype(v)
	})
}

// UpdateMimetype sets the "mimetype" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateMimetype() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateMimetype()
	})
}

// SetWidth sets the "width" field.
func (u *AssetUpsertBulk) SetWidth(v int) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetWidth(v)
	})
}

// AddWidth adds v to the "width" field.
func (u *AssetUpsertBulk) AddWidth(v int) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.AddWidth(v)
	})
}

// UpdateWidth sets the "width" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateWidth() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateWidth()
	})
}

// SetHeight sets the "height" field.
func (u *AssetUpsertBulk) SetHeight(v int) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetHeight(v)
	})
}

// AddHeight adds v to the "height" field.
func (u *AssetUpsertBulk) AddHeight(v int) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.AddHeight(v)
	})
}

// UpdateHeight sets the "height" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateHeight() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateHeight()
	})
}

// SetAccountID sets the "account_id" field.
func (u *AssetUpsertBulk) SetAccountID(v xid.ID) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateAccountID() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateAccountID()
	})
}

// Exec executes the query.
func (u *AssetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AssetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AssetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AssetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
