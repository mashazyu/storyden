// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/link"
	"github.com/Southclaws/storyden/internal/ent/node"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/Southclaws/storyden/internal/ent/tag"
	"github.com/rs/xid"
)

// NodeQuery is the builder for querying Node entities.
type NodeQuery struct {
	config
	ctx        *QueryContext
	order      []node.OrderOption
	inters     []Interceptor
	predicates []predicate.Node
	withOwner  *AccountQuery
	withParent *NodeQuery
	withNodes  *NodeQuery
	withAssets *AssetQuery
	withTags   *TagQuery
	withLinks  *LinkQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NodeQuery builder.
func (nq *NodeQuery) Where(ps ...predicate.Node) *NodeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NodeQuery) Limit(limit int) *NodeQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NodeQuery) Offset(offset int) *NodeQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NodeQuery) Unique(unique bool) *NodeQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NodeQuery) Order(o ...node.OrderOption) *NodeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryOwner chains the current query on the "owner" edge.
func (nq *NodeQuery) QueryOwner() *AccountQuery {
	query := (&AccountClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, node.OwnerTable, node.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (nq *NodeQuery) QueryParent() *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, node.ParentTable, node.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodes chains the current query on the "nodes" edge.
func (nq *NodeQuery) QueryNodes() *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, node.NodesTable, node.NodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssets chains the current query on the "assets" edge.
func (nq *NodeQuery) QueryAssets() *AssetQuery {
	query := (&AssetClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, node.AssetsTable, node.AssetsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (nq *NodeQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, node.TagsTable, node.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinks chains the current query on the "links" edge.
func (nq *NodeQuery) QueryLinks() *LinkQuery {
	query := (&LinkClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, selector),
			sqlgraph.To(link.Table, link.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, node.LinksTable, node.LinksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Node entity from the query.
// Returns a *NotFoundError when no Node was found.
func (nq *NodeQuery) First(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{node.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NodeQuery) FirstX(ctx context.Context) *Node {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Node ID from the query.
// Returns a *NotFoundError when no Node ID was found.
func (nq *NodeQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{node.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NodeQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Node entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Node entity is found.
// Returns a *NotFoundError when no Node entities are found.
func (nq *NodeQuery) Only(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{node.Label}
	default:
		return nil, &NotSingularError{node.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NodeQuery) OnlyX(ctx context.Context) *Node {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Node ID in the query.
// Returns a *NotSingularError when more than one Node ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NodeQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{node.Label}
	default:
		err = &NotSingularError{node.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NodeQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nodes.
func (nq *NodeQuery) All(ctx context.Context) ([]*Node, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Node, *NodeQuery]()
	return withInterceptors[[]*Node](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NodeQuery) AllX(ctx context.Context) []*Node {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Node IDs.
func (nq *NodeQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(node.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NodeQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NodeQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NodeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NodeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NodeQuery) Clone() *NodeQuery {
	if nq == nil {
		return nil
	}
	return &NodeQuery{
		config:     nq.config,
		ctx:        nq.ctx.Clone(),
		order:      append([]node.OrderOption{}, nq.order...),
		inters:     append([]Interceptor{}, nq.inters...),
		predicates: append([]predicate.Node{}, nq.predicates...),
		withOwner:  nq.withOwner.Clone(),
		withParent: nq.withParent.Clone(),
		withNodes:  nq.withNodes.Clone(),
		withAssets: nq.withAssets.Clone(),
		withTags:   nq.withTags.Clone(),
		withLinks:  nq.withLinks.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithOwner(opts ...func(*AccountQuery)) *NodeQuery {
	query := (&AccountClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withOwner = query
	return nq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithParent(opts ...func(*NodeQuery)) *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withParent = query
	return nq
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithNodes(opts ...func(*NodeQuery)) *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNodes = query
	return nq
}

// WithAssets tells the query-builder to eager-load the nodes that are connected to
// the "assets" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithAssets(opts ...func(*AssetQuery)) *NodeQuery {
	query := (&AssetClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withAssets = query
	return nq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithTags(opts ...func(*TagQuery)) *NodeQuery {
	query := (&TagClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withTags = query
	return nq
}

// WithLinks tells the query-builder to eager-load the nodes that are connected to
// the "links" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithLinks(opts ...func(*LinkQuery)) *NodeQuery {
	query := (&LinkClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withLinks = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Node.Query().
//		GroupBy(node.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NodeQuery) GroupBy(field string, fields ...string) *NodeGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NodeGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = node.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Node.Query().
//		Select(node.FieldCreatedAt).
//		Scan(ctx, &v)
func (nq *NodeQuery) Select(fields ...string) *NodeSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NodeSelect{NodeQuery: nq}
	sbuild.label = node.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NodeSelect configured with the given aggregations.
func (nq *NodeQuery) Aggregate(fns ...AggregateFunc) *NodeSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !node.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Node, error) {
	var (
		nodes       = []*Node{}
		_spec       = nq.querySpec()
		loadedTypes = [6]bool{
			nq.withOwner != nil,
			nq.withParent != nil,
			nq.withNodes != nil,
			nq.withAssets != nil,
			nq.withTags != nil,
			nq.withLinks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Node).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Node{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withOwner; query != nil {
		if err := nq.loadOwner(ctx, query, nodes, nil,
			func(n *Node, e *Account) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withParent; query != nil {
		if err := nq.loadParent(ctx, query, nodes, nil,
			func(n *Node, e *Node) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNodes; query != nil {
		if err := nq.loadNodes(ctx, query, nodes,
			func(n *Node) { n.Edges.Nodes = []*Node{} },
			func(n *Node, e *Node) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withAssets; query != nil {
		if err := nq.loadAssets(ctx, query, nodes,
			func(n *Node) { n.Edges.Assets = []*Asset{} },
			func(n *Node, e *Asset) { n.Edges.Assets = append(n.Edges.Assets, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withTags; query != nil {
		if err := nq.loadTags(ctx, query, nodes,
			func(n *Node) { n.Edges.Tags = []*Tag{} },
			func(n *Node, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withLinks; query != nil {
		if err := nq.loadLinks(ctx, query, nodes,
			func(n *Node) { n.Edges.Links = []*Link{} },
			func(n *Node, e *Link) { n.Edges.Links = append(n.Edges.Links, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NodeQuery) loadOwner(ctx context.Context, query *AccountQuery, nodes []*Node, init func(*Node), assign func(*Node, *Account)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Node)
	for i := range nodes {
		fk := nodes[i].AccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nq *NodeQuery) loadParent(ctx context.Context, query *NodeQuery, nodes []*Node, init func(*Node), assign func(*Node, *Node)) error {
	ids := make([]xid.ID, 0, len(nodes))
	nodeids := make(map[xid.ID][]*Node)
	for i := range nodes {
		fk := nodes[i].ParentNodeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(node.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_node_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nq *NodeQuery) loadNodes(ctx context.Context, query *NodeQuery, nodes []*Node, init func(*Node), assign func(*Node, *Node)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[xid.ID]*Node)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(node.FieldParentNodeID)
	}
	query.Where(predicate.Node(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(node.NodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentNodeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parent_node_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nq *NodeQuery) loadAssets(ctx context.Context, query *AssetQuery, nodes []*Node, init func(*Node), assign func(*Node, *Asset)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Node)
	nids := make(map[xid.ID]map[*Node]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(node.AssetsTable)
		s.Join(joinT).On(s.C(asset.FieldID), joinT.C(node.AssetsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(node.AssetsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(node.AssetsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Node]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Asset](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "assets" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nq *NodeQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Node, init func(*Node), assign func(*Node, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Node)
	nids := make(map[xid.ID]map[*Node]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(node.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(node.TagsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(node.TagsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(node.TagsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Node]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nq *NodeQuery) loadLinks(ctx context.Context, query *LinkQuery, nodes []*Node, init func(*Node), assign func(*Node, *Link)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[xid.ID]*Node)
	nids := make(map[xid.ID]map[*Node]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(node.LinksTable)
		s.Join(joinT).On(s.C(link.FieldID), joinT.C(node.LinksPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(node.LinksPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(node.LinksPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(xid.ID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*xid.ID)
				inValue := *values[1].(*xid.ID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Node]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Link](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "links" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (nq *NodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeString))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for i := range fields {
			if fields[i] != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if nq.withOwner != nil {
			_spec.Node.AddColumnOnce(node.FieldAccountID)
		}
		if nq.withParent != nil {
			_spec.Node.AddColumnOnce(node.FieldParentNodeID)
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(node.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = node.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range nq.modifiers {
		m(selector)
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (nq *NodeQuery) Modify(modifiers ...func(s *sql.Selector)) *NodeSelect {
	nq.modifiers = append(nq.modifiers, modifiers...)
	return nq.Select()
}

// NodeGroupBy is the group-by builder for Node entities.
type NodeGroupBy struct {
	selector
	build *NodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NodeGroupBy) Aggregate(fns ...AggregateFunc) *NodeGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeQuery, *NodeGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NodeGroupBy) sqlScan(ctx context.Context, root *NodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NodeSelect is the builder for selecting fields of Node entities.
type NodeSelect struct {
	*NodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NodeSelect) Aggregate(fns ...AggregateFunc) *NodeSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeQuery, *NodeSelect](ctx, ns.NodeQuery, ns, ns.inters, v)
}

func (ns *NodeSelect) sqlScan(ctx context.Context, root *NodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ns *NodeSelect) Modify(modifiers ...func(s *sql.Selector)) *NodeSelect {
	ns.modifiers = append(ns.modifiers, modifiers...)
	return ns
}