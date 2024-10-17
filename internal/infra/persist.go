package infra

import (
	"github.com/Sreeram-ganesan/my-blog/internal/adapters/persist"
	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"github.com/Sreeram-ganesan/my-blog/internal/core/di"
	"github.com/Sreeram-ganesan/my-blog/internal/core/outport"
)

func wirePersistPorts(
	cfg *app.Config,
	cache outport.Cache,
	di *di.DI,
) func() {
	pers := persist.NewPersistence(cfg)
	addrBook := persist.NewAddrBookAdapter(
		pers,
		cache,
	)
	di.UseCases.AddrBook = addrBook
	return pers.Close
}
