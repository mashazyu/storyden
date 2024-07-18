package index_job

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"go.uber.org/zap"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/library"
	"github.com/Southclaws/storyden/app/resources/mq"
	"github.com/Southclaws/storyden/app/resources/post"
	"github.com/Southclaws/storyden/app/resources/post/reply"
	"github.com/Southclaws/storyden/app/resources/profile"
	"github.com/Southclaws/storyden/app/services/semdex"
	"github.com/Southclaws/storyden/internal/pubsub"
)

type indexerConsumer struct {
	l *zap.Logger

	replyRepo   reply.Repository
	nodeRepo    library.Repository
	accountRepo account.Repository

	qnode    pubsub.Topic[mq.IndexNode]
	qnodesum pubsub.Topic[mq.SummariseNode]
	qpost    pubsub.Topic[mq.IndexPost]

	indexer   semdex.Indexer
	retriever semdex.Retriever
}

func newIndexConsumer(
	l *zap.Logger,

	replyRepo reply.Repository,
	nodeRepo library.Repository,
	accountRepo account.Repository,

	qnode pubsub.Topic[mq.IndexNode],
	qnodesum pubsub.Topic[mq.SummariseNode],
	qpost pubsub.Topic[mq.IndexPost],
	qprofile pubsub.Topic[mq.IndexProfile],

	indexer semdex.Indexer,
	retriever semdex.Retriever,
) *indexerConsumer {
	return &indexerConsumer{
		l:           l,
		replyRepo:   replyRepo,
		nodeRepo:    nodeRepo,
		accountRepo: accountRepo,
		qnode:       qnode,
		qnodesum:    qnodesum,
		qpost:       qpost,
		indexer:     indexer,
		retriever:   retriever,
	}
}

func (i *indexerConsumer) indexPost(ctx context.Context, id post.ID) error {
	p, err := i.replyRepo.Get(ctx, id)
	if err != nil {
		return fault.Wrap(err, fctx.With(ctx))
	}

	return i.indexer.Index(ctx, p)
}

func (i *indexerConsumer) indexNode(ctx context.Context, id library.NodeID) error {
	n, err := i.nodeRepo.GetByID(ctx, id)
	if err != nil {
		return fault.Wrap(err, fctx.With(ctx))
	}

	err = i.indexer.Index(ctx, n)
	if err != nil {
		return fault.Wrap(err, fctx.With(ctx))
	}

	err = i.qnodesum.Publish(ctx, mq.SummariseNode{ID: n.ID})
	if err != nil {
		return fault.Wrap(err, fctx.With(ctx))
	}

	return nil
}

func (i *indexerConsumer) indexProfile(ctx context.Context, id account.AccountID) error {
	p, err := i.accountRepo.GetByID(ctx, id)
	if err != nil {
		return fault.Wrap(err, fctx.With(ctx))
	}

	return i.indexer.Index(ctx, profile.ProfileFromAccount(p))
}
