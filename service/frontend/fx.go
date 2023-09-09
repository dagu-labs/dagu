package frontend

import (
	"context"
	"github.com/dagu-dev/dagu/internal/config"
	"github.com/dagu-dev/dagu/internal/logger"
	"github.com/dagu-dev/dagu/service/frontend/handlers"
	"github.com/dagu-dev/dagu/service/frontend/http"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(handlers.NewDAG, fx.ResultTags(`group:"handlers"`))),
	fx.Provide(New),
)

type Params struct {
	fx.In

	Config   *config.Config
	Logger   logger.Logger
	Handlers []http.Handler `group:"handlers"`
}

func LifetimeHooks(lc fx.Lifecycle, srv *http.Server) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				return srv.Serve(ctx)
			},
			OnStop: func(_ context.Context) error {
				srv.Shutdown()
				return nil
			},
		},
	)
}

func New(params Params) *http.Server {
	serverParams := http.ServerParams{
		Host:     params.Config.Host,
		Port:     params.Config.Port,
		TLS:      params.Config.TLS,
		Logger:   params.Logger,
		Handlers: params.Handlers,
	}

	if params.Config.IsBasicAuth {
		serverParams.BasicAuth = &http.BasicAuth{
			Username: params.Config.BasicAuthUsername,
			Password: params.Config.BasicAuthUsername,
		}
	}

	return http.NewServer(serverParams)
}
