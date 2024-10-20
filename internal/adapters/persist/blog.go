package persist

import (
	"context"

	"github.com/Sreeram-ganesan/my-blog/internal/adapters/persist/internal/mapper"
	"github.com/Sreeram-ganesan/my-blog/internal/adapters/persist/internal/repo"
	"github.com/Sreeram-ganesan/my-blog/internal/core/model"
)

type blogBookAdapter struct {
	repo *repo.BlogBookRepo
	// blogByIdCache cache.blogByIdCache
}

func (a *blogBookAdapter) AddBlog(ctx context.Context, b *model.BlogToSave) (*model.Blog, error) {
	entity := mapper.BlogToSaveModelToEntity(b)
	entity, err := a.repo.AddBlog(ctx, entity)
	if err != nil {
		return nil, err
	}
	blog := mapper.BlogEntityToModel(entity)
	// a.blogByIdCache.Set(ctx, blog)
	return blog, nil
}
