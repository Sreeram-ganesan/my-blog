package usecase

import (
	"context"

	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"github.com/Sreeram-ganesan/my-blog/internal/core/model"
)

func (uc *UseCases) AddBlog(
	ctx context.Context,
	blog *model.BlogToSave,
) (*model.Blog, error) {
	app.Logger(ctx).Debugf("Add blog: %v", blog)
	newBlog, err := uc.BlogBook.AddBlog(ctx, blog)
	if err != nil {
		app.Logger(ctx).Errorf("Adding new blog failed with error: %v", err)
		return nil, err
	}
	app.Logger(ctx).Debugf("Added blog: %v", newBlog)
	return newBlog, nil
}
