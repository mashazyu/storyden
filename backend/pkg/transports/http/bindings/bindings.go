package bindings

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/Southclaws/storyden/backend/internal/config"
	"github.com/Southclaws/storyden/backend/internal/web"
	"github.com/Southclaws/storyden/backend/pkg/transports/http/openapi"
)

// Bindings is a DI parameter struct that is used to compose together all of the
// individual service bindings in this package. When the provider below depends
// on this type, it provides all these composed bindings to the DI system so the
// invoke call can mount them onto the router using the OpenAPI ServerInterface.
type Bindings struct {
	fx.In
	Authentication
}

func mountBindings(lc fx.Lifecycle, l *zap.Logger, router chi.Router, si openapi.ServerInterface) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			openapi.HandlerFromMux(si, router)

			// quick and simple version handler
			router.Get("/version", func(w http.ResponseWriter, _ *http.Request) {
				web.Write(w, map[string]string{"version": config.Version}) //nolint:errcheck
			})

			return nil
		},
	})

	l.Info("mounted OpenAPI to service bindings")
}

func addMiddleware(l *zap.Logger, router chi.Router, a Authentication) {
	router.Use(web.WithLogger)
	router.Use(a.middleware)

	l.Info("added router middleware")
}

func Build() fx.Option {
	return fx.Options(
		// Provide the bindings struct which implements the generated OpenAPI
		// interface by composing together all of the service bindings into a
		// single struct.
		fx.Provide(func(s Bindings) openapi.ServerInterface { return &s }),

		// Add the middleware bindings.
		fx.Invoke(addMiddleware),

		// Mount the bound OpenAPI routes onto the router.
		fx.Invoke(mountBindings),

		// Provide all service layer bindings to the DI system so they can be
		// depended upon during the binding provider above.
		fx.Provide(
			NewAuthentication,
		),
	)
}