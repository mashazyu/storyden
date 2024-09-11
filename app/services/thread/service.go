// Package thread provides APIs for working with threads which are sequences of
// posts. Threads can be created with one post, listed, searched and updated.
package thread

import (
	"context"
	"net/url"

	"github.com/Southclaws/opt"
	"github.com/rs/xid"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/account/account_querier"
	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/datagraph/semdex"
	"github.com/Southclaws/storyden/app/resources/mq"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/resources/post/category"
	"github.com/Southclaws/storyden/app/resources/post/thread"
	"github.com/Southclaws/storyden/app/resources/rbac"
	"github.com/Southclaws/storyden/app/resources/visibility"
	"github.com/Southclaws/storyden/app/services/link/fetcher"
	"github.com/Southclaws/storyden/app/services/mention/mentioner"
	"github.com/Southclaws/storyden/internal/infrastructure/pubsub"
)

type Service interface {
	// Create a new thread in the specified category.
	Create(
		ctx context.Context,
		title string,
		authorID account.AccountID,
		categoryID category.CategoryID,
		status visibility.Visibility,
		tags []string,
		meta map[string]any,
		partial Partial,
	) (*thread.Thread, error)

	Update(ctx context.Context, threadID post.ID, partial Partial) (*thread.Thread, error)

	Delete(ctx context.Context, id post.ID) error

	List(ctx context.Context,
		page int,
		size int,
		opts Params,
	) (*thread.Result, error)

	// Get one thread and the posts within it.
	Get(
		ctx context.Context,
		threadID post.ID,
	) (*thread.Thread, error)
}

type Partial struct {
	Title      opt.Optional[string]
	Content    opt.Optional[datagraph.Content]
	Tags       opt.Optional[[]xid.ID]
	Category   opt.Optional[xid.ID]
	Visibility opt.Optional[visibility.Visibility]
	URL        opt.Optional[url.URL]
	Meta       opt.Optional[map[string]any]
}

func (p Partial) Opts() (opts []thread.Option) {
	p.Title.Call(func(v string) { opts = append(opts, thread.WithTitle(v)) })
	p.Content.Call(func(v datagraph.Content) { opts = append(opts, thread.WithContent(v)) })
	p.Tags.Call(func(v []xid.ID) { opts = append(opts, thread.WithTags(v)) })
	p.Category.Call(func(v xid.ID) { opts = append(opts, thread.WithCategory(xid.ID(v))) })
	p.Visibility.Call(func(v visibility.Visibility) { opts = append(opts, thread.WithVisibility(v)) })
	p.Meta.Call(func(v map[string]any) { opts = append(opts, thread.WithMeta(v)) })
	return
}

func Build() fx.Option {
	return fx.Provide(New)
}

type service struct {
	l    *zap.Logger
	rbac rbac.AccessManager

	accountQuery account_querier.Querier
	thread_repo  thread.Repository
	fetcher      *fetcher.Fetcher
	recommender  semdex.Recommender
	indexQueue   pubsub.Topic[mq.IndexPost]
	mentioner    *mentioner.Mentioner
}

func New(
	l *zap.Logger,
	rbac rbac.AccessManager,

	accountQuery account_querier.Querier,
	thread_repo thread.Repository,
	fetcher *fetcher.Fetcher,
	recommender semdex.Recommender,
	indexQueue pubsub.Topic[mq.IndexPost],
	mentioner *mentioner.Mentioner,
) Service {
	return &service{
		l:            l.With(zap.String("service", "thread")),
		rbac:         rbac,
		accountQuery: accountQuery,
		thread_repo:  thread_repo,
		fetcher:      fetcher,
		recommender:  recommender,
		indexQueue:   indexQueue,
		mentioner:    mentioner,
	}
}
