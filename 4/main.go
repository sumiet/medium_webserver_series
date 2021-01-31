package main

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	httpServer "medium/medium_webserver_series/4/http"
	"medium/medium_webserver_series/4/loggerfx"
	rpcServer "medium/medium_webserver_series/4/rpc"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Provide(rpcServer.New),
		fx.Invoke(httpServer.New),
		fx.Invoke(registerHooks),
		loggerfx.Module,
	).Run()
}

func registerHooks(
	lifecycle fx.Lifecycle, mux *http.ServeMux, logger *zap.SugaredLogger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				// start the rpc server
				l, err := net.Listen("tcp", ":8081")
				logger.Errorf("Error while starting rpc server: %+v", err)
				go func() {
					for {
						rpc.Accept(l)
					}
				}()
				logger.Info("Listening on port 8081 for RPC requests")


				// start the http server
				logger.Info("Listening on localhost:8080 for HTTP requests")
				go http.ListenAndServe(":8080", mux) // we will look into how to gracefully handle these errors later

				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
