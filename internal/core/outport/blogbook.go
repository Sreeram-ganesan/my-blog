package outport

import (
	"context"

	"github.com/Sreeram-ganesan/my-blog/internal/core/model"
)

type BlogBook interface {
	AddBlog(ctx context.Context, c *model.BlogToSave) (*model.Blog, error)
}
