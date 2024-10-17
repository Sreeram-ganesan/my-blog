package infra

import (
	"context"
	"go.uber.org/zap"
	"github.com/Sreeram-ganesan/my-blog/internal/adapters/apiserver"
	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
)

func Start(deployment string) {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	ctx := app.ContextWithLogger(context.Background(), zap.S())

	cfg := app.LoadConfig(deployment)
	di := wireDependencies(cfg)
	apiserver.Start(ctx, di)
}
