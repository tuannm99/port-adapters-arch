package edgeconfig

import "context"

type Service struct {
	renderer Renderer
}

func NewService(renderer Renderer) *Service {
	return &Service{renderer: renderer}
}

func (s *Service) BuildWebsiteConfig(
	ctx context.Context,
	domain string,
	upstream string,
) (string, error) {
	return s.renderer.Render(ctx, RenderInput{
		Domain:   domain,
		Upstream: upstream,
	})
}
