// Code generated by ent, DO NOT EDIT.

package collection

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the collection type in the database.
	Label = "collection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the collection in the database.
	Table = "collections"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "collections"
	// OwnerInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	OwnerInverseTable = "accounts"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "account_collections"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "collection_posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
)

// Columns holds all SQL columns for collection fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldVisibility,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "collections"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_collections",
}

var (
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"collection_id", "post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityDraft is the default value of the Visibility enum.
const DefaultVisibility = VisibilityDraft

// Visibility values.
const (
	VisibilityDraft  Visibility = "draft"
	VisibilityReview Visibility = "review"
	VisibilityPublic Visibility = "public"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityDraft, VisibilityReview, VisibilityPublic:
		return nil
	default:
		return fmt.Errorf("collection: invalid enum value for visibility field: %q", v)
	}
}

// OrderOption defines the ordering options for the Collection queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PostsTable, PostsPrimaryKey...),
	)
}
