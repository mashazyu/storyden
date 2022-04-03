// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: "(No description)"},
		{Name: "colour", Type: field.TypeString, Default: "#8577ce"},
		{Name: "sort", Type: field.TypeInt, Default: -1},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "post_category", Type: field.TypeString, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_posts_category",
				Columns:    []*schema.Column{CategoriesColumns[6]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "read", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "subscription_id", Type: field.TypeString},
		{Name: "subscription_notifications", Type: field.TypeString, Nullable: true},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "notifications_subscriptions_notifications",
				Columns:    []*schema.Column{NotificationsColumns[7]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "slug", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString},
		{Name: "short", Type: field.TypeString},
		{Name: "first", Type: field.TypeBool},
		{Name: "pinned", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "root_post_id", Type: field.TypeString, Nullable: true},
		{Name: "reply_post_id", Type: field.TypeString, Nullable: true},
		{Name: "category_id", Type: field.TypeString, Nullable: true},
		{Name: "category_posts", Type: field.TypeString, Nullable: true},
		{Name: "post_posts", Type: field.TypeString, Nullable: true},
		{Name: "react_post", Type: field.TypeString, Nullable: true},
		{Name: "tag_posts", Type: field.TypeString, Nullable: true},
		{Name: "user_posts", Type: field.TypeUUID, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_categories_posts",
				Columns:    []*schema.Column{PostsColumns[14]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_posts_posts",
				Columns:    []*schema.Column{PostsColumns[15]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_reacts_Post",
				Columns:    []*schema.Column{PostsColumns[16]},
				RefColumns: []*schema.Column{ReactsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_tags_posts",
				Columns:    []*schema.Column{PostsColumns[17]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[18]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReactsColumns holds the columns for the "reacts" table.
	ReactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "emoji", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "post_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "post_reacts", Type: field.TypeString, Nullable: true},
		{Name: "user_reacts", Type: field.TypeUUID, Nullable: true},
	}
	// ReactsTable holds the schema information for the "reacts" table.
	ReactsTable = &schema.Table{
		Name:       "reacts",
		Columns:    ReactsColumns,
		PrimaryKey: []*schema.Column{ReactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reacts_posts_reacts",
				Columns:    []*schema.Column{ReactsColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reacts_users_reacts",
				Columns:    []*schema.Column{ReactsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RulesColumns holds the columns for the "rules" table.
	RulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "server_id", Type: field.TypeString, Nullable: true},
		{Name: "server_ru", Type: field.TypeString, Nullable: true},
	}
	// RulesTable holds the schema information for the "rules" table.
	RulesTable = &schema.Table{
		Name:       "rules",
		Columns:    RulesColumns,
		PrimaryKey: []*schema.Column{RulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rules_servers_ru",
				Columns:    []*schema.Column{RulesColumns[4]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ServersColumns holds the columns for the "servers" table.
	ServersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "ip", Type: field.TypeString},
		{Name: "hn", Type: field.TypeString},
		{Name: "pc", Type: field.TypeInt},
		{Name: "pm", Type: field.TypeInt},
		{Name: "gm", Type: field.TypeString},
		{Name: "la", Type: field.TypeString},
		{Name: "pa", Type: field.TypeBool},
		{Name: "vn", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "banner", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
		{Name: "active", Type: field.TypeBool},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "rule_server", Type: field.TypeInt, Nullable: true},
	}
	// ServersTable holds the schema information for the "servers" table.
	ServersTable = &schema.Table{
		Name:       "servers",
		Columns:    ServersColumns,
		PrimaryKey: []*schema.Column{ServersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "servers_rules_Server",
				Columns:    []*schema.Column{ServersColumns[16]},
				RefColumns: []*schema.Column{RulesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "refers_type", Type: field.TypeEnum, Enums: []string{"FORUM_POST_RESPONSE"}},
		{Name: "refers_to", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "notification_subscription", Type: field.TypeString, Nullable: true},
		{Name: "user_subscriptions", Type: field.TypeUUID, Nullable: true},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscriptions_notifications_subscription",
				Columns:    []*schema.Column{SubscriptionsColumns[7]},
				RefColumns: []*schema.Column{NotificationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "subscriptions_users_subscriptions",
				Columns:    []*schema.Column{SubscriptionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "post_tags", Type: field.TypeString, Nullable: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_posts_tags",
				Columns:    []*schema.Column{TagsColumns[2]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "post_author", Type: field.TypeString, Nullable: true},
		{Name: "react_user", Type: field.TypeString, Nullable: true},
		{Name: "server_user", Type: field.TypeString, Nullable: true},
		{Name: "subscription_user", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_posts_author",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_reacts_user",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{ReactsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_servers_User",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{ServersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_subscriptions_user",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{SubscriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostReplyToColumns holds the columns for the "post_replyTo" table.
	PostReplyToColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeString},
		{Name: "replyTo_id", Type: field.TypeString},
	}
	// PostReplyToTable holds the schema information for the "post_replyTo" table.
	PostReplyToTable = &schema.Table{
		Name:       "post_replyTo",
		Columns:    PostReplyToColumns,
		PrimaryKey: []*schema.Column{PostReplyToColumns[0], PostReplyToColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_replyTo_post_id",
				Columns:    []*schema.Column{PostReplyToColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_replyTo_replyTo_id",
				Columns:    []*schema.Column{PostReplyToColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostRepliesColumns holds the columns for the "post_replies" table.
	PostRepliesColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeString},
		{Name: "reply_id", Type: field.TypeString},
	}
	// PostRepliesTable holds the schema information for the "post_replies" table.
	PostRepliesTable = &schema.Table{
		Name:       "post_replies",
		Columns:    PostRepliesColumns,
		PrimaryKey: []*schema.Column{PostRepliesColumns[0], PostRepliesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_replies_post_id",
				Columns:    []*schema.Column{PostRepliesColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_replies_reply_id",
				Columns:    []*schema.Column{PostRepliesColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		NotificationsTable,
		PostsTable,
		ReactsTable,
		RulesTable,
		ServersTable,
		SubscriptionsTable,
		TagsTable,
		UsersTable,
		PostReplyToTable,
		PostRepliesTable,
	}
)

func init() {
	CategoriesTable.ForeignKeys[0].RefTable = PostsTable
	NotificationsTable.ForeignKeys[0].RefTable = SubscriptionsTable
	PostsTable.ForeignKeys[0].RefTable = CategoriesTable
	PostsTable.ForeignKeys[1].RefTable = PostsTable
	PostsTable.ForeignKeys[2].RefTable = ReactsTable
	PostsTable.ForeignKeys[3].RefTable = TagsTable
	PostsTable.ForeignKeys[4].RefTable = UsersTable
	ReactsTable.ForeignKeys[0].RefTable = PostsTable
	ReactsTable.ForeignKeys[1].RefTable = UsersTable
	RulesTable.ForeignKeys[0].RefTable = ServersTable
	ServersTable.ForeignKeys[0].RefTable = RulesTable
	SubscriptionsTable.ForeignKeys[0].RefTable = NotificationsTable
	SubscriptionsTable.ForeignKeys[1].RefTable = UsersTable
	TagsTable.ForeignKeys[0].RefTable = PostsTable
	UsersTable.ForeignKeys[0].RefTable = PostsTable
	UsersTable.ForeignKeys[1].RefTable = ReactsTable
	UsersTable.ForeignKeys[2].RefTable = ServersTable
	UsersTable.ForeignKeys[3].RefTable = SubscriptionsTable
	PostReplyToTable.ForeignKeys[0].RefTable = PostsTable
	PostReplyToTable.ForeignKeys[1].RefTable = PostsTable
	PostRepliesTable.ForeignKeys[0].RefTable = PostsTable
	PostRepliesTable.ForeignKeys[1].RefTable = PostsTable
}