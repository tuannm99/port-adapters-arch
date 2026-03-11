package websiteapp

import (
	"context"

	"github.com/tuannm99/edge-platform/apps/control-plane/internal/domain/website"
)

type Repository interface {
	Create(ctx context.Context, w website.Website) error
	List(ctx context.Context) ([]website.Website, error)
	GetByID(ctx context.Context, id string) (website.Website, error)
}
