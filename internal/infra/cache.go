package infra

import (
	"fmt"
	"github.com/Sreeram-ganesan/my-blog/internal/adapters/cache"
	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"github.com/Sreeram-ganesan/my-blog/internal/core/di"
	"github.com/Sreeram-ganesan/my-blog/internal/core/outport"
)

func wireCachePorts(cfg *app.Config, _ *di.DI) (outport.Cache, func()) {
	switch cfg.Cache.Type {
	case "none":
		return cache.NewNoCache(), func() {}
	case "inmem":
		return cache.NewInMemCache(), func() {}
	case "redis":
		r := cache.NewRedisCache(&cfg.Cache.Redis)
		return r, func() {
			r.Close()
		}
	default:
		panic(fmt.Sprintf("unknown cache type: %s", cfg.Cache.Type))
	}
}
