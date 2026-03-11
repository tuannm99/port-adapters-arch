package websiteapp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/tuannm99/edge-platform/apps/control-plane/internal/domain/website"
)

//	type ConfigGenerator interface {
//		Generate(cfg nginxconf.ServerConfig) (string, error)
//	}

type Repository interface {
	Create(ctx context.Context, w website.Website) error
	List(ctx context.Context) ([]website.Website, error)
	GetByID(ctx context.Context, id string) (website.Website, error)
}

type CreateInput struct {
	Domain   string `json:"domain"`
	Upstream string `json:"upstream"`
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, in CreateInput) (website.Website, error) {
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

func (s *Service) List(ctx context.Context) ([]website.Website, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (website.Website, error) {
	return s.repo.GetByID(ctx, id)
}
