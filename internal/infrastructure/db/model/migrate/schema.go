// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// AuthenticationsColumns holds the columns for the "authentications" table.
	AuthenticationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "service", Type: field.TypeString},
		{Name: "identifier", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "account_authentication", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// AuthenticationsTable holds the schema information for the "authentications" table.
	AuthenticationsTable = &schema.Table{
		Name:       "authentications",
		Columns:    AuthenticationsColumns,
		PrimaryKey: []*schema.Column{AuthenticationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "authentications_accounts_authentication",
				Columns:    []*schema.Column{AuthenticationsColumns[6]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "authentication_service_identifier",
				Unique:  true,
				Columns: []*schema.Column{AuthenticationsColumns[2], AuthenticationsColumns[3]},
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: "(No description)"},
		{Name: "colour", Type: field.TypeString, Default: "#8577ce"},
		{Name: "sort", Type: field.TypeInt, Default: -1},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "read", Type: field.TypeBool},
		{Name: "notification_subscription", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "subscription_notifications", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notifications_subscriptions_subscription",
				Columns:    []*schema.Column{NotificationsColumns[6]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "notifications_subscriptions_notifications",
				Columns:    []*schema.Column{NotificationsColumns[7]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "first", Type: field.TypeBool},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "slug", Type: field.TypeString, Nullable: true},
		{Name: "pinned", Type: field.TypeBool, Default: false},
		{Name: "body", Type: field.TypeString},
		{Name: "short", Type: field.TypeString},
		{Name: "account_posts", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "category_id", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "root_post_id", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "reply_to_post_id", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_accounts_posts",
				Columns:    []*schema.Column{PostsColumns[10]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "posts_categories_posts",
				Columns:    []*schema.Column{PostsColumns[11]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_posts_posts",
				Columns:    []*schema.Column{PostsColumns[12]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_posts_replies",
				Columns:    []*schema.Column{PostsColumns[13]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReactsColumns holds the columns for the "reacts" table.
	ReactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "emoji", Type: field.TypeString},
		{Name: "account_reacts", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "post_reacts", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "react_account", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "react_post", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// ReactsTable holds the schema information for the "reacts" table.
	ReactsTable = &schema.Table{
		Name:       "reacts",
		Columns:    ReactsColumns,
		PrimaryKey: []*schema.Column{ReactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reacts_accounts_reacts",
				Columns:    []*schema.Column{ReactsColumns[3]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reacts_posts_reacts",
				Columns:    []*schema.Column{ReactsColumns[4]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reacts_accounts_account",
				Columns:    []*schema.Column{ReactsColumns[5]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reacts_posts_Post",
				Columns:    []*schema.Column{ReactsColumns[6]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "refers_type", Type: field.TypeString},
		{Name: "refers_to", Type: field.TypeString},
		{Name: "account_subscriptions", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "subscription_account", Type: field.TypeBytes, Nullable: true, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscriptions_accounts_subscriptions",
				Columns:    []*schema.Column{SubscriptionsColumns[4]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "subscriptions_accounts_account",
				Columns:    []*schema.Column{SubscriptionsColumns[5]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// RoleAccountsColumns holds the columns for the "role_accounts" table.
	RoleAccountsColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "account_id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// RoleAccountsTable holds the schema information for the "role_accounts" table.
	RoleAccountsTable = &schema.Table{
		Name:       "role_accounts",
		Columns:    RoleAccountsColumns,
		PrimaryKey: []*schema.Column{RoleAccountsColumns[0], RoleAccountsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_accounts_role_id",
				Columns:    []*schema.Column{RoleAccountsColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_accounts_account_id",
				Columns:    []*schema.Column{RoleAccountsColumns[1]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagPostsColumns holds the columns for the "tag_posts" table.
	TagPostsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
		{Name: "post_id", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{"mysql": "binary(12)", "postgres": "bytea"}},
	}
	// TagPostsTable holds the schema information for the "tag_posts" table.
	TagPostsTable = &schema.Table{
		Name:       "tag_posts",
		Columns:    TagPostsColumns,
		PrimaryKey: []*schema.Column{TagPostsColumns[0], TagPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_posts_tag_id",
				Columns:    []*schema.Column{TagPostsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_posts_post_id",
				Columns:    []*schema.Column{TagPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AuthenticationsTable,
		CategoriesTable,
		NotificationsTable,
		PostsTable,
		ReactsTable,
		RolesTable,
		SubscriptionsTable,
		TagsTable,
		RoleAccountsTable,
		TagPostsTable,
	}
)

func init() {
	AuthenticationsTable.ForeignKeys[0].RefTable = AccountsTable
	NotificationsTable.ForeignKeys[0].RefTable = SubscriptionsTable
	NotificationsTable.ForeignKeys[1].RefTable = SubscriptionsTable
	PostsTable.ForeignKeys[0].RefTable = AccountsTable
	PostsTable.ForeignKeys[1].RefTable = CategoriesTable
	PostsTable.ForeignKeys[2].RefTable = PostsTable
	PostsTable.ForeignKeys[3].RefTable = PostsTable
	ReactsTable.ForeignKeys[0].RefTable = AccountsTable
	ReactsTable.ForeignKeys[1].RefTable = PostsTable
	ReactsTable.ForeignKeys[2].RefTable = AccountsTable
	ReactsTable.ForeignKeys[3].RefTable = PostsTable
	SubscriptionsTable.ForeignKeys[0].RefTable = AccountsTable
	SubscriptionsTable.ForeignKeys[1].RefTable = AccountsTable
	RoleAccountsTable.ForeignKeys[0].RefTable = RolesTable
	RoleAccountsTable.ForeignKeys[1].RefTable = AccountsTable
	TagPostsTable.ForeignKeys[0].RefTable = TagsTable
	TagPostsTable.ForeignKeys[1].RefTable = PostsTable
}
