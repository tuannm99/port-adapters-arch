package websiteapp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/domain/website"
)

var _ WebsiteUseCase = (*WebsiteService)(nil)

type WebsiteUseCase interface {
	Create(ctx context.Context, in CreateInput) (website.Website, error)
	List(ctx context.Context) ([]website.Website, error)
	GetByID(ctx context.Context, id string) (website.Website, error)
}

type CreateInput struct {
	Domain   string
	Upstream string
}

type WebsiteService struct {
	repo Repository
}

func NewWebsiteService(repo Repository) *WebsiteService {
	return &WebsiteService{repo: repo}
}

func (s *WebsiteService) Create(ctx context.Context, in CreateInput) (website.Website, error) {
	if in.Domain == "" {
		return website.Website{}, fmt.Errorf("domain is required")
	}
	if in.Upstream == "" {
		return website.Website{}, fmt.Errorf("upstream is required")
	}

	w := website.Website{
		ID:       uuid.NewString(),
		Domain:   in.Domain,
		Upstream: in.Upstream,
	}

	if err := s.repo.Create(ctx, w); err != nil {
		return website.Website{}, err
	}

	return w, nil
}

func (s *WebsiteService) List(ctx context.Context) ([]website.Website, error) {
	return s.repo.List(ctx)
}

func (s *WebsiteService) GetByID(ctx context.Context, id string) (website.Website, error) {
	return s.repo.GetByID(ctx, id)
}
