// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/HarryBird/url-shorten/app/gateway/internal/data"
	"github.com/HarryBird/url-shorten/app/gateway/internal/server"
	"github.com/HarryBird/url-shorten/app/gateway/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, app *conf.App, tracerProvider *trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewRedis(confData, logger)
	discovery := server.NewDiscover(registry, app, logger)
	shortenClient := data.NewShortenServiceClient(discovery, tracerProvider, logger)
	dataData, cleanup, err := data.NewData(client, shortenClient, logger)
	if err != nil {
		return nil, nil, err
	}
	shortenRepo := data.NewShortenRepo(dataData, logger)
	shortenCase := biz.NewShortenCase(app, shortenRepo, logger)
	gatewayService := service.NewGatewayService(shortenCase, logger)
	httpServer := server.NewHTTPServer(confServer, tracerProvider, gatewayService, logger)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, gatewayService, logger)
	registrar := server.NewRegistrar(registry, app, logger)
	kratosApp := newApp(logger, httpServer, grpcServer, registrar)
	return kratosApp, func() {
		cleanup()
	}, nil
}
