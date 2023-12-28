// Package semdex provides an interface for semantic indexing of the datagraph.
package semdex

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/kr/pretty"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources/datagraph"
)

type Service interface {
	Index(ctx context.Context, object datagraph.Indexable) error
}

type noop struct{}

func (n noop) Index(ctx context.Context, object datagraph.Indexable) error {
	return nil
}

func New(lc fx.Lifecycle, wc *weaviate.Client) (Service, error) {
	if wc == nil {
		return noop{}, nil
	}

	lc.Append(fx.StartHook(func(ctx context.Context) error {
		classObj := &models.Class{
			Class:      "Content",
			Vectorizer: "text2vec-openai",
			ModuleConfig: map[string]interface{}{
				"text2vec-openai":   map[string]interface{}{},
				"generative-openai": map[string]interface{}{},
			},
		}

		r, err := wc.Schema().
			ClassExistenceChecker().
			WithClassName("Content").
			Do(ctx)
		if err != nil {
			return fault.Wrap(err)
		}

		if !r {
			err := wc.Schema().
				ClassCreator().
				WithClass(classObj).
				Do(ctx)
			if err != nil {
				return fault.Wrap(err)
			}
		}

		return nil
	}))

	return &service{wc}, nil
}

type service struct {
	wc *weaviate.Client
}

func (s *service) Index(ctx context.Context, object datagraph.Indexable) error {
	response, err := s.wc.Data().Creator().
		WithClassName("Content").
		WithProperties(map[string]any{
			"content": object.Text(),
			"props":   object.Props(),
		}).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	_ = response

	pretty.Println(response)

	return nil
}