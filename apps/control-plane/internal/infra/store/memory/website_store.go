package memory

import (
	"context"
	"fmt"
	"sync"

	websiteapp "github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/application/website"
	"github.com/tuannm99/port-adapters-arch/apps/control-plane/internal/domain/website"
)

var _ websiteapp.Repository = (*WebsiteStore)(nil)

type WebsiteStore struct {
	mu       sync.RWMutex
	websites []website.Website
}

func NewWebsiteStore() *WebsiteStore {
	return &WebsiteStore{
		websites: make([]website.Website, 0),
	}
}

func (s *WebsiteStore) Create(_ context.Context, w website.Website) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.websites = append(s.websites, w)
	return nil
}

func (s *WebsiteStore) List(_ context.Context) ([]website.Website, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]website.Website, len(s.websites))
	copy(out, s.websites)
	return out, nil
}

func (s *WebsiteStore) GetByID(_ context.Context, id string) (website.Website, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, item := range s.websites {
		if item.ID == id {
			return item, nil
		}
	}

	return website.Website{}, fmt.Errorf("website not found")
}
