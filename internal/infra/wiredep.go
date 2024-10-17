package infra

import (
	"go.uber.org/zap"
	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"github.com/Sreeram-ganesan/my-blog/internal/core/di"
	"github.com/Sreeram-ganesan/my-blog/internal/core/usecase"
)

func wireDependencies(cfg *app.Config) *di.DI {
	zap.S().Info("Initialize DI objects")
	newDI := &di.DI{
		Config:   cfg,
		UseCases: &usecase.UseCases{},
	}

	cache, cacheCleanup := wireCachePorts(cfg, newDI)

	persistCleanup := wirePersistPorts(
		cfg,
		cache,
		newDI,
	)

	newDI.Close = func() {
		zap.S().Info("Performing cleanup of all initialized DI objects")
		persistCleanup()
		cacheCleanup()
	}
	return newDI
}
