// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/category"
	"github.com/Southclaws/storyden/internal/ent/link"
	"github.com/Southclaws/storyden/internal/ent/post"
	"github.com/rs/xid"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// First holds the value of the "first" field.
	First bool `json:"first,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Pinned holds the value of the "pinned" field.
	Pinned bool `json:"pinned,omitempty"`
	// RootPostID holds the value of the "root_post_id" field.
	RootPostID xid.ID `json:"root_post_id,omitempty"`
	// ReplyToPostID holds the value of the "reply_to_post_id" field.
	ReplyToPostID xid.ID `json:"reply_to_post_id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Short holds the value of the "short" field.
	Short string `json:"short,omitempty"`
	// Arbitrary metadata used by clients to store domain specific information.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Visibility holds the value of the "visibility" field.
	Visibility post.Visibility `json:"visibility,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID xid.ID `json:"category_id,omitempty"`
	// LinkID holds the value of the "link_id" field.
	LinkID xid.ID `json:"link_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges         PostEdges `json:"edges"`
	account_posts *xid.ID
	selectValues  sql.SelectValues
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Author holds the value of the author edge.
	Author *Account `json:"author,omitempty"`
	// Category is only required for root posts. It should never be added to a child post.
	Category *Category `json:"category,omitempty"`
	// Tags are only required for root posts. It should never be added to a child post.
	Tags []*Tag `json:"tags,omitempty"`
	// A many-to-many recursive self reference. The root post is the first post in the thread.
	Root *Post `json:"root,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// A many-to-many recursive self reference. The replyTo post is an optional post that this post is in reply to.
	ReplyTo *Post `json:"replyTo,omitempty"`
	// Replies holds the value of the replies edge.
	Replies []*Post `json:"replies,omitempty"`
	// Reacts holds the value of the reacts edge.
	Reacts []*React `json:"reacts,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*LikePost `json:"likes,omitempty"`
	// Mentions holds the value of the mentions edge.
	Mentions []*MentionProfile `json:"mentions,omitempty"`
	// Assets holds the value of the assets edge.
	Assets []*Asset `json:"assets,omitempty"`
	// Collections holds the value of the collections edge.
	Collections []*Collection `json:"collections,omitempty"`
	// Link holds the value of the link edge.
	Link *Link `json:"link,omitempty"`
	// ContentLinks holds the value of the content_links edge.
	ContentLinks []*Link `json:"content_links,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [14]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) AuthorOrErr() (*Account, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) CategoryOrErr() (*Category, error) {
	if e.Category != nil {
		return e.Category, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: category.Label}
	}
	return nil, &NotLoadedError{edge: "category"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// RootOrErr returns the Root value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) RootOrErr() (*Post, error) {
	if e.Root != nil {
		return e.Root, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: post.Label}
	}
	return nil, &NotLoadedError{edge: "root"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[4] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// ReplyToOrErr returns the ReplyTo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) ReplyToOrErr() (*Post, error) {
	if e.ReplyTo != nil {
		return e.ReplyTo, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: post.Label}
	}
	return nil, &NotLoadedError{edge: "replyTo"}
}

// RepliesOrErr returns the Replies value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) RepliesOrErr() ([]*Post, error) {
	if e.loadedTypes[6] {
		return e.Replies, nil
	}
	return nil, &NotLoadedError{edge: "replies"}
}

// ReactsOrErr returns the Reacts value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) ReactsOrErr() ([]*React, error) {
	if e.loadedTypes[7] {
		return e.Reacts, nil
	}
	return nil, &NotLoadedError{edge: "reacts"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) LikesOrErr() ([]*LikePost, error) {
	if e.loadedTypes[8] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// MentionsOrErr returns the Mentions value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) MentionsOrErr() ([]*MentionProfile, error) {
	if e.loadedTypes[9] {
		return e.Mentions, nil
	}
	return nil, &NotLoadedError{edge: "mentions"}
}

// AssetsOrErr returns the Assets value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) AssetsOrErr() ([]*Asset, error) {
	if e.loadedTypes[10] {
		return e.Assets, nil
	}
	return nil, &NotLoadedError{edge: "assets"}
}

// CollectionsOrErr returns the Collections value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CollectionsOrErr() ([]*Collection, error) {
	if e.loadedTypes[11] {
		return e.Collections, nil
	}
	return nil, &NotLoadedError{edge: "collections"}
}

// LinkOrErr returns the Link value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) LinkOrErr() (*Link, error) {
	if e.Link != nil {
		return e.Link, nil
	} else if e.loadedTypes[12] {
		return nil, &NotFoundError{label: link.Label}
	}
	return nil, &NotLoadedError{edge: "link"}
}

// ContentLinksOrErr returns the ContentLinks value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) ContentLinksOrErr() ([]*Link, error) {
	if e.loadedTypes[13] {
		return e.ContentLinks, nil
	}
	return nil, &NotLoadedError{edge: "content_links"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldMetadata:
			values[i] = new([]byte)
		case post.FieldFirst, post.FieldPinned:
			values[i] = new(sql.NullBool)
		case post.FieldTitle, post.FieldSlug, post.FieldBody, post.FieldShort, post.FieldVisibility:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt, post.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case post.FieldID, post.FieldRootPostID, post.FieldReplyToPostID, post.FieldCategoryID, post.FieldLinkID:
			values[i] = new(xid.ID)
		case post.ForeignKeys[0]: // account_posts
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				po.ID = *value
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case post.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				po.DeletedAt = new(time.Time)
				*po.DeletedAt = value.Time
			}
		case post.FieldFirst:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field first", values[i])
			} else if value.Valid {
				po.First = value.Bool
			}
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case post.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				po.Slug = value.String
			}
		case post.FieldPinned:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pinned", values[i])
			} else if value.Valid {
				po.Pinned = value.Bool
			}
		case post.FieldRootPostID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field root_post_id", values[i])
			} else if value != nil {
				po.RootPostID = *value
			}
		case post.FieldReplyToPostID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field reply_to_post_id", values[i])
			} else if value != nil {
				po.ReplyToPostID = *value
			}
		case post.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				po.Body = value.String
			}
		case post.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				po.Short = value.String
			}
		case post.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &po.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case post.FieldVisibility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				po.Visibility = post.Visibility(value.String)
			}
		case post.FieldCategoryID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value != nil {
				po.CategoryID = *value
			}
		case post.FieldLinkID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field link_id", values[i])
			} else if value != nil {
				po.LinkID = *value
			}
		case post.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_posts", values[i])
			} else if value.Valid {
				po.account_posts = new(xid.ID)
				*po.account_posts = *value.S.(*xid.ID)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Post.
// This includes values selected through modifiers, order, etc.
func (po *Post) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the Post entity.
func (po *Post) QueryAuthor() *AccountQuery {
	return NewPostClient(po.config).QueryAuthor(po)
}

// QueryCategory queries the "category" edge of the Post entity.
func (po *Post) QueryCategory() *CategoryQuery {
	return NewPostClient(po.config).QueryCategory(po)
}

// QueryTags queries the "tags" edge of the Post entity.
func (po *Post) QueryTags() *TagQuery {
	return NewPostClient(po.config).QueryTags(po)
}

// QueryRoot queries the "root" edge of the Post entity.
func (po *Post) QueryRoot() *PostQuery {
	return NewPostClient(po.config).QueryRoot(po)
}

// QueryPosts queries the "posts" edge of the Post entity.
func (po *Post) QueryPosts() *PostQuery {
	return NewPostClient(po.config).QueryPosts(po)
}

// QueryReplyTo queries the "replyTo" edge of the Post entity.
func (po *Post) QueryReplyTo() *PostQuery {
	return NewPostClient(po.config).QueryReplyTo(po)
}

// QueryReplies queries the "replies" edge of the Post entity.
func (po *Post) QueryReplies() *PostQuery {
	return NewPostClient(po.config).QueryReplies(po)
}

// QueryReacts queries the "reacts" edge of the Post entity.
func (po *Post) QueryReacts() *ReactQuery {
	return NewPostClient(po.config).QueryReacts(po)
}

// QueryLikes queries the "likes" edge of the Post entity.
func (po *Post) QueryLikes() *LikePostQuery {
	return NewPostClient(po.config).QueryLikes(po)
}

// QueryMentions queries the "mentions" edge of the Post entity.
func (po *Post) QueryMentions() *MentionProfileQuery {
	return NewPostClient(po.config).QueryMentions(po)
}

// QueryAssets queries the "assets" edge of the Post entity.
func (po *Post) QueryAssets() *AssetQuery {
	return NewPostClient(po.config).QueryAssets(po)
}

// QueryCollections queries the "collections" edge of the Post entity.
func (po *Post) QueryCollections() *CollectionQuery {
	return NewPostClient(po.config).QueryCollections(po)
}

// QueryLink queries the "link" edge of the Post entity.
func (po *Post) QueryLink() *LinkQuery {
	return NewPostClient(po.config).QueryLink(po)
}

// QueryContentLinks queries the "content_links" edge of the Post entity.
func (po *Post) QueryContentLinks() *LinkQuery {
	return NewPostClient(po.config).QueryContentLinks(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := po.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("first=")
	builder.WriteString(fmt.Sprintf("%v", po.First))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(po.Title)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(po.Slug)
	builder.WriteString(", ")
	builder.WriteString("pinned=")
	builder.WriteString(fmt.Sprintf("%v", po.Pinned))
	builder.WriteString(", ")
	builder.WriteString("root_post_id=")
	builder.WriteString(fmt.Sprintf("%v", po.RootPostID))
	builder.WriteString(", ")
	builder.WriteString("reply_to_post_id=")
	builder.WriteString(fmt.Sprintf("%v", po.ReplyToPostID))
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(po.Body)
	builder.WriteString(", ")
	builder.WriteString("short=")
	builder.WriteString(po.Short)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", po.Metadata))
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", po.Visibility))
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", po.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("link_id=")
	builder.WriteString(fmt.Sprintf("%v", po.LinkID))
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post
