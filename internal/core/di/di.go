package di

import (
	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"github.com/Sreeram-ganesan/my-blog/internal/core/usecase"
)

type DI struct {
	Close    func()
	Config   *app.Config
	UseCases *usecase.UseCases
}
