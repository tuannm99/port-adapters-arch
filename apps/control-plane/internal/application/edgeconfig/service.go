package edgeconfig

import "context"

var _ EdgeConfigUseCase = (*EdgeService)(nil)

type EdgeConfigUseCase interface {
	BuildWebsiteConfig(ctx context.Context, domain, upstream string) (string, error)
}

type EdgeService struct {
	renderer Renderer
}

func NewEdgeService(renderer Renderer) *EdgeService {
	return &EdgeService{renderer: renderer}
}

func (s *EdgeService) BuildWebsiteConfig(
	ctx context.Context,
	domain string,
	upstream string,
) (string, error) {
	return s.renderer.Render(ctx, RenderInput{
		Domain:   domain,
		Upstream: upstream,
	})
}
